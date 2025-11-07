package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	//"google.golang.org/genai"

	"log"
	"net/http"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type Comment struct {
	ID       uint   `gorm:"primaryKey"`
	Date     string `gorm:"not null"`
	UUID     string `gorm:"not null"`
	Note     string `gorm:"not null"`
	UserNote string `gorm:"not null"`
}

func getPrompt(uuid string, date string, name string, note string, chr string) *string {
	// Status取得
	status := retGetStatus(uuid, date)
	//strStatus := fmt.Sprintf("Mood: %v, Enjoyment: %v", status.Mood, status.Enjoyment)
	//fmt.Println(strStatus)

	subjectID := retGetSubjectByUserID(uuid)

	todos := retGetTODOByUserID(uuid, date)
	//fmt.Println("ToDos:", todos)

	strToDos := ""
	for _, v := range todos {
		var finished string
		if v.Checked {
			finished = "完了しました"
		} else {
			finished = "完了してません"
		}

		todo := fmt.Sprintf(" - ToDoのタイトルは%vです。このタスクは%v。このタスクの優先順位をは%vです}\n", v.Title, finished, v.Status)
		strToDos += todo
	}

	//fmt.Println(strToDos)

	logs := retGetLogByUserID(uuid, date)
	//fmt.Println("StudyLogs:", logs)
	sum := 0
	for _, slog := range logs {
		sum += slog.StudyTime
	}

	var logstr string = ""
	for _, slog := range logs {
		str := fmt.Sprintf(" - 科目のID: %v、勉強開始時刻:%v:%v、終了時刻: %v:%v \n", slog.SubjectID, slog.StartHours, slog.StartMinutes, slog.EndHours, slog.EndMinutes, slog.StudyTime)
		logstr += str
	}
	//fmt.Println(logstr)

	unfinished := retGetUnfinishedByUserID(uuid)
	//fmt.Println("Unfinished:", unfinished)

	strUnfinished := ""

	for _, list := range unfinished {
		strUnfinished += fmt.Sprintf(" - タイトル:%v¥、優先順位:%v、日付:%v、科目ID:%v\n", list.Title, list.Status, list.Date, list.SubjectID)
	}

	mood := status.Mood
	enjoyment := status.Enjoyment
	studyTime := fmt.Sprintf("%d時間%d分", sum/60, sum%60)
	todolist := strToDos
	strTrack := logstr

	feature := "### キャラの性格を示すプロンプト\n"
	if chr == "先生" {
		feature = `あなたは、優しく頼れる先生です。やさしく論理的に接してください。
私のことを後から示す呼び名で「さん」をつけて呼んでください。
 私が勉強を続けられるように、具体的な称賛と提案を提供します。
あなたのコメントは、私が前向きに勉強を続けるための大きな助けとなります。
常に伴奏姿勢で指導してください。中高生にも伝わる言葉を用い、命令形よりも「〜してみましょう」を優先してください。
`
	}
	if chr == "友達" {
		feature = `あなたは、気持ちに寄り添う親友です。フランクで等身大に接してください。
私のことを後から示す呼び名で呼んでください。
まず、気持ちを言い換えて共感してください
丁寧語をつかわず、タメ口で話してください。
`

	}
	if chr == "親" {
		feature = `あなたは、温かく見守る親です。包み込む安心感で接してください。
私のことを後から示す呼び名で呼んでください。
まず、気持ちを言い換えて共感してください。
丁寧語ではなく、優しい口調で話してください。
`
	}

	prompt := fmt.Sprintf(feature+`あなたは、私の%sです。
私の呼び名は%sです。
私が勉強をがんばれるような声かけをするのが得意です。

### 以下のデータだけを事実として用い、事実と異なる推測はしないでください
 - 気分スコアは4段階で「4が最高」「1が最低」「0は記入なし」
 - ToDoのCheckedがtrueは完了、falseは未完了
 - 「今日の楽しみ」が空文字のときは触れない
 - 「勉強時間」はすでに合計済みなので再計算しない
 - JSONは正として扱い、矛盾があればJSONを優先する
 - 180〜220字程度
 - 誇張や一般化は避け、データに即した具体的な称賛と提案にする

### 次の順で200字程度のコメントを書いてください（段落は分けない）:

今日は%sです
なお、返事は来ないものとしてマークダウンや箇条書きで番号リストではなく、先生が返すような文章のみで作成してください。

## 添付情報
### 今日の気分（4段階）(0は記入なし)
%d
### 今日の楽しみ(空文字は記入なし)
%s
### 勉強時間（注意：勉強時間について言及する際はここだけを見るようにしてください）
%s
### 全科目ID一覧(ToDoやStudyLog。未完了リストのSubjectIDに対応)
%s
### ToDoの達成状況（Checked: trueは完了タスク、Checked: falseは未完了タスクです。これらを踏まえた上で言及してください）
%s
### TimeTracking
%s
### 未完了リスト
%s
### 今日の振り返り
%s
`, chr, name, date, mood, enjoyment, studyTime, subjectID, todolist, strTrack, strUnfinished, note)
	fmt.Println(prompt)
	lenNote := len(note)
	lenPrompt := len(prompt)

	fmt.Printf("Length of note: %d, Length of prompt: %d\n", lenNote, lenPrompt)

	if lenNote > 4000 {
		fmt.Println("Warning: Note length exceeds 4000 characters, which may cause issues with token limits.")
		return nil
	}

	if lenPrompt > 6000 {
		fmt.Println("Warning: Prompt length exceeds 6000 characters, which may cause issues with token limits.")
		return nil
	}

	return &prompt
}

func reqComment(c *gin.Context) {
	uuid := GetProfile(c).UUID

	type RequestBody struct {
		Date string `json:"Date"`
		Note string `json:"Note"`
		Chr  string `json:"Chr"`
	}
	var req RequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var name string = "あなた"
	yourProfile := GetProfile(c)
	if yourProfile != nil {
		name = yourProfile.Username
	}

	ctx := context.Background()

	// OpenAI version
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPEN_API_KEY is not set")
	}

	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	prompt := getPrompt(uuid, req.Date, name, req.Note, req.Chr)
	if prompt == nil {
		c.JSON(500, gin.H{"error": "AIに送信するデータの量が多すぎます。コメントを短くしてください。"})
		return
	}

	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(*prompt),
		},
		Model: openai.ChatModelGPT4_1Nano,
	})
	defer stream.Close()

	acc := openai.ChatCompletionAccumulator{}

	response := ""
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	flusher, _ := c.Writer.(http.Flusher)

	for stream.Next() {
		chunk := stream.Current()
		acc.AddChunk(chunk)
		if len(chunk.Choices) > 0 {
			char := chunk.Choices[0].Delta.Content
			fmt.Print(char) // 届いた分だけ出力
			response += char
			fmt.Fprintf(c.Writer, char)
			flusher.Flush()
		}
	}
	if err := stream.Err(); err != nil {
		log.Fatal(err)
		fmt.Printf("failed to generate content: %w", err)
		c.JSON(500, gin.H{"error": "コンテンツの生成に失敗しました"})
		return
	}

	saveUserComment(yourProfile.UUID, req.Date, req.Note)
	saveAIComment(uuid, req.Date, response)

	fmt.Println("Success generating content")

	// GeminiAPI version

	//apiKey := os.Getenv("GEMINI_API_KEY")
	//apiKey := os.Getenv("OPENAI_API_KEY")
	//if apiKey == "" {
	//	//log.Fatal("GOOGLE_API_KEY is not set")
	//}

	//client, err := genai.NewClient(ctx, &genai.ClientConfig{
	//	Backend: genai.BackendGeminiAPI, // AI StudioのGemini APIを使う
	//	APIKey:  apiKey,
	//})
	//
	//if err != nil {
	//	fmt.Printf("failed to create client: %v", err)
	//	c.JSON(500, gin.H{"error": "AIクライアントの作成に失敗しました"})
	//	return
	//}
	//
	//modelName := "gemini-2.5-flash"
	//
	//prompt := getPrompt(uuid, req.Date, name, req.Note, req.Chr)
	//
	//contents := genai.Text(prompt)
	//
	//response := ""
	//
	//c.Writer.Header().Set("Content-Type", "text/event-stream")
	//flusher, _ := c.Writer.(http.Flusher)
	//
	//for resp, err := range client.Models.GenerateContentStream(ctx, modelName, contents, nil) {
	//	if err != nil {
	//		fmt.Printf("failed to generate content: %w", err)
	//		c.JSON(500, gin.H{"error": "コンテンツの生成に失敗しました"})
	//		return
	//	}
	//
	//	chunk := resp.Text()
	//	//fmt.Println(chunk)
	//	response += chunk
	//	fmt.Fprintf(c.Writer, "%s", chunk)
	//	flusher.Flush()
	//}

	//saveAIComment(uuid, req.Date, response)
	//
	//fmt.Println("Success generating content")
}

func retGetComment(uuid string, date string) *Comment {
	var comment Comment
	err := db.Model(&Comment{}).Where("date = ? AND uuid = ?", date, uuid).First(&comment).Error
	//fmt.Println(comment)
	if err != nil {
		fmt.Println("Error fetching comment:", err)
		return nil
	}

	return &comment
}

func getComment(c *gin.Context) {
	fmt.Println("comment/")
	uuid := GetProfile(c).UUID
	date := c.Query("date")

	comment := retGetComment(uuid, date)
	fmt.Println(comment)

	c.JSON(http.StatusOK, comment)
}

func saveAIComment(uuid string, date string, note string) {
	comment := Comment{
		Date: date,
		UUID: uuid,
		Note: note,
	}
	if retGetComment(uuid, date) == nil {
		fmt.Println("No existing comment, creating new one.")
		if err := db.Create(&comment).Error; err != nil {
			fmt.Println("Error creating comment:", err)
			return
		}
		fmt.Println("Success creating comment")
	} else {
		if err := db.Model(&Comment{}).Where("date = ? AND uuid = ?", date, uuid).Updates(comment).Error; err != nil {
			fmt.Println("Error updating comment:", err)
			return
		}
		fmt.Println("Success updating comment")
	}
}

func saveUserComment(uuid string, date string, note string) {
	comment := Comment{
		Date:     date,
		UUID:     uuid,
		UserNote: note,
	}
	if retGetComment(uuid, date) == nil {
		fmt.Println("No existing comment, creating new one.")
		if err := db.Create(&comment).Error; err != nil {
			fmt.Println("Error creating comment:", err)
			return
		}
		fmt.Println("Success creating comment")
	} else {
		if err := db.Model(&Comment{}).Where("date = ? AND uuid = ?", date, uuid).Updates(comment).Error; err != nil {
			fmt.Println("Error updating comment:", err)
			return
		}
		fmt.Println("Success updating comment")
	}
}
