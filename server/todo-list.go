package main

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type TODOLIST struct {
	ID        uint       `gorm:"primaryKey"`
	Date      string     `gorm:"not null"`
	SubjectID int        `gorm:"not null"`
	UUID      string     `gorm:"not null"`
	Title     string     `gorm:"not null"`
	Checked   bool       `gorm:"not null"`
	Status    string     `gorm:"not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	CheckedAt *time.Time
}

func retGetTODOByUserID(uuid string, date string) []TODOLIST {
	var todos []TODOLIST

	res := db.Model(&TODOLIST{}).Where("uuid = ? and date = ?", uuid, date).Find(&todos)
	if res.Error != nil {
		fmt.Println("Error fetching ToDo List:", res.Error)
		return nil
	}
	return todos
}

func retGetTODO_Checked(uuid string, date string) []TODOLIST {
	var todos []TODOLIST

	res := db.Model(&TODOLIST{}).Where("uuid = ? and date = ? and checked = 1", uuid, date).Find(&todos)
	if res.Error != nil {
		fmt.Println("Error fetching ToDo List:", res.Error)
		return nil
	}
	return todos
}

func retGetTODObyStatusUnfinished(uuid string, date string, status string) []TODOLIST {
	var todos []TODOLIST

	res := db.Model(&TODOLIST{}).Where(`uuid = ? and date = ? and checked = 0 and status = '`+status+`'`, uuid, date).Find(&todos)
	if res.Error != nil {
		fmt.Println("Error fetching ToDo List:", res.Error)
		return nil
	}
	return todos
}

func getTODOByUserID(c *gin.Context) {
	fmt.Println("todo/")
	uuid := GetProfile(c).UUID

	date := c.Query("date")

	todos := retGetTODOByUserID(uuid, date)
	if todos == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}

	fmt.Println("Fetched ToDo:", len(todos))
	c.JSON(http.StatusOK, todos)
}

func getToDOByGroup(c *gin.Context) {
	fmt.Println("todo/group")
	uuid := GetProfile(c).UUID

	date := c.Query("date")

	checked := retGetTODO_Checked(uuid, date)
	if checked == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched ToDo:", len(checked))

	MUST := retGetTODObyStatusUnfinished(uuid, date, "MUST")
	if MUST == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched ToDo:", len(MUST))

	WANT := retGetTODObyStatusUnfinished(uuid, date, "WANT")
	if WANT == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched ToDo:", len(WANT))

	c.JSON(http.StatusOK, gin.H{"checked": checked, "MUST": MUST, "WANT": WANT})
}

func AddToDo(c *gin.Context) {
	fmt.Println("todo/add")

	//body, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(string(body))
	//
	uuid := GetProfile(c).UUID
	var req struct {
		Date      string `json:"date"`
		Title     string `json:"title"`
		Status    string `json:"status"`
		SubjectID int    `json:"subjectID"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	todo := TODOLIST{
		Date:      req.Date,
		SubjectID: req.SubjectID,
		UUID:      uuid,
		Title:     req.Title,
		Checked:   false,
		Status:    req.Status,
		CheckedAt: nil,
	}

	if err := db.Model(&TODOLIST{}).Create(&todo).Error; err != nil {
		fmt.Println("Error creating ToDo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの作成に失敗しました"})
		return
	}
	fmt.Println("Sccuess creating ToDO")
	c.JSON(http.StatusOK, gin.H{"message": "ToDoを作成しました", "id": todo.ID})
}

func ToDoChecked(c *gin.Context) {
	fmt.Println("todo/checked")
	var req struct {
		ID int `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	var todo TODOLIST
	if err := db.Model(&TODOLIST{}).Where("id = ?", req.ID).First(&todo).Error; err != nil {
		fmt.Println("Error fetching ToDo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの取得に失敗しました"})
		return
	}

	todo.Checked = !todo.Checked
	now := time.Now()
	if todo.Checked {
		todo.CheckedAt = &now
	} else {
		todo.CheckedAt = nil
	}

	if err := db.Save(&todo).Error; err != nil {
		fmt.Println("Error updating ToDo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの更新に失敗しました"})
		return
	}

	fmt.Println("Sccuess updating ToDO")
	c.JSON(http.StatusOK, gin.H{"message": "ToDoを更新しました"})
}

func ToDoEdit(c *gin.Context) {
	fmt.Println("todo/edit")
	var req struct {
		ID       int    `json:"id"`
		NewTitle string `json:"newtitle"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}
	var todo TODOLIST
	if err := db.Model(&TODOLIST{}).Where("id = ?", req.ID).
		First(&todo).Error; err != nil {
		fmt.Println("Error fetching ToDo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの取得に失敗しました"})
		return
	}
	todo.Title = req.NewTitle

	if todo.Title == "" {
		if err := db.Model(&TODOLIST{}).Where("id = ?", req.ID).Delete(&TODOLIST{}).Error; err != nil {
			fmt.Println("Error deleting ToDo:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの削除に失敗しました"})
			return
		}
		fmt.Println("ToDo deleted successfully")
	} else {
		if err := db.Save(&todo).Error; err != nil {
			fmt.Println("Error updating ToDo:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの更新に失敗しました"})
			return
		}
		fmt.Println("Sccuess updating ToDO")
	}
	c.JSON(http.StatusOK, gin.H{"message": "ToDoを更新しました"})
}

type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
}

type Notify struct {
	ID          uint      `gorm:"primaryKey"`
	DateTime    time.Time `gorm:"not null"`
	ToDoID      int       `gorm:"not null"`
	UUID        string    `gorm:"not null"`
	Endpoint    string    `gorm:"not null"`
	P256dh      string    `gorm:"not null"`
	Auth        string    `gorm:"not null"`
	CreatedDate string    `gorm:"not null"`
}

// func push(when time.Time, c *gin.Context, uuid string, nTaskid int, Endpoint string, P256dh string, Auth string, vapidPrivateKey string, vapidPublicKey string) {
func push(notify Notify, c *gin.Context, vapidPrivateKey string, vapidPublicKey string) {
	d := time.Until(notify.DateTime)
	if d <= 0 {
		fmt.Println("⏱️ もう過ぎてるので実行しません:", notify.DateTime)
		return
	}

	// ゴルーチンで裏実行
	go func() {
		fmt.Println("✅ スケジュール登録:", notify.DateTime)
		time.Sleep(d)
		dNotify(notify.ID)
		fmt.Println("🟢 実行しました！:", time.Now())

		var todos TODOLIST

		res := db.Model(&TODOLIST{}).Where(`uuid = ? and id = ?`, notify.UUID, notify.ToDoID).First(&todos)
		if res.Error != nil {
			fmt.Println("Error fetching ToDo List:", res.Error)
		}
		fmt.Println("Fetched ToDo Title:", todos.Title)
		titleOfTask := todos.Title

		// 通知内容
		message := map[string]string{
			"title": "時間になりました",
			"body":  titleOfTask,
		}
		payload, _ := json.Marshal(message)
		log.Printf("📝 Payload: %s", string(payload))

		// WebPush送信
		log.Println("🚀 Sending notification...")
		resp, err := webpush.SendNotification(payload, &webpush.Subscription{
			Endpoint: notify.Endpoint,
			Keys: webpush.Keys{
				P256dh: notify.P256dh,
				Auth:   notify.Auth,
			},
		}, &webpush.Options{
			VAPIDPrivateKey: vapidPrivateKey,
			VAPIDPublicKey:  vapidPublicKey,
			TTL:             30,
			Subscriber:      "mailto:test@example.com",
		})

		if err != nil {
			log.Printf("❌ WebPush Send Error: %v", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		responseBody, _ := io.ReadAll(resp.Body)
		log.Printf("✅ WebPush sent successfully!")
		log.Printf("📊 Status Code: %d", resp.StatusCode)
		log.Printf("📄 Response: %s", string(responseBody))
	}()
}

func dNotify(id uint) {
	var n Notify
	n.ID = id
	if err := db.Delete(&n).Error; err != nil {
		fmt.Println("Error deleting Notify:", err)
		return
	}
	fmt.Println("Success deleting Notify")
}

type NotifyResponse struct {
	ID       uint      `json:"id"`
	DateTime time.Time `json:"date_time"`
	ToDoID   int       `json:"to_do_id"`
}

func retGetNotifyByUUID(uuid string, datetime string) []NotifyResponse {
	//var notifies []Notify

	var notifies []NotifyResponse

	res := db.Model(&Notify{}).Select("id", "date_time", "to_do_id").Where("uuid = ? AND created_date = ?", uuid, datetime).Scan(&notifies)
	if res.Error != nil {
		fmt.Println("Error fetching Notify:", res.Error)
		return nil
	}
	if res.RowsAffected == 0 {
		return []NotifyResponse{}
	}
	return notifies
}

func getNotifyByUUID(c *gin.Context) {
	fmt.Println("api/notify/")
	uuid := GetProfile(c).UUID
	datetime := c.Query("datetime")

	notifies := retGetNotifyByUUID(uuid, datetime)
	if notifies == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notifyの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched Notify:", len(notifies))
	c.JSON(http.StatusOK, notifies)
}

func refreshNotify() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// VAPID鍵の検証
	vapidPrivateKey := os.Getenv("VAPID_PRIVATE_KEY")
	vapidPublicKey := os.Getenv("VAPID_PUBLIC_KEY")

	var notifies []Notify

	res := db.Model(&Notify{}).Find(&notifies)
	if res.Error != nil {
		fmt.Println("Error fetching Notify:", res.Error)
	}
	if res.RowsAffected == 0 {
		return
	}
	for _, notify := range notifies {
		push(notify, nil, vapidPrivateKey, vapidPublicKey)
	}
}

func notify(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// VAPID鍵の検証
	vapidPrivateKey := os.Getenv("VAPID_PRIVATE_KEY")
	vapidPublicKey := os.Getenv("VAPID_PUBLIC_KEY")

	log.Printf("=== VAPID Keys Check ===")
	log.Printf("Private Key set: %v (length: %d)", vapidPrivateKey != "", len(vapidPrivateKey))
	log.Printf("Public Key set: %v (length: %d)", vapidPublicKey != "", len(vapidPublicKey))

	if vapidPrivateKey == "" || vapidPublicKey == "" {
		log.Fatal("❌ Error: VAPID keys not set in environment variables")
	}

	// 鍵の形式を確認（秘密鍵は短く、公開鍵は長い）
	if len(vapidPrivateKey) > len(vapidPublicKey) {
		log.Fatal("❌ Error: VAPID keys appear to be swapped!")
	}

	fmt.Println("/api/send")

	datetimeTemp := c.GetHeader("datetime")
	nTaskid := c.GetHeader("Task")
	createdDate := c.GetHeader("CreatedDate")
	fmt.Println((c.GetHeader("Authorization")))
	uuid := GetProfile(c).UUID
	fmt.Println(uuid)

	fmt.Println(datetimeTemp)
	fmt.Println(nTaskid)

	// ① その時刻を「日本時間(Asia/Tokyo)」として解釈
	loc, _ := time.LoadLocation("Asia/Tokyo")
	layout := "2006-01-02T15:04" // 秒なし
	datetime, err := time.ParseInLocation(layout, datetimeTemp, loc)
	if err != nil {
		log.Printf("❌ Datetime Parse Error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid datetime format"})
		return
	}

	fmt.Println("Tokyo:", datetime) // 2025-11-10 21:47:00 +0900 JST

	log.Println("=== /api/send endpoint called ===")

	var sub Subscription
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("❌ Error reading body: %v", err)
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}

	log.Printf("📦 Received body: %s", string(body))

	if err := json.Unmarshal(body, &sub); err != nil {
		log.Printf("❌ JSON Unmarshal Error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	log.Printf("📍 Endpoint: %s", sub.Endpoint)
	log.Printf("🔑 P256dh length: %d", len(sub.Keys.P256dh))
	log.Printf("🔑 Auth length: %d", len(sub.Keys.Auth))

	//loc, _ := time.LoadLocation("Asia/Tokyo")

	// 例①：特定日時で
	runAt := datetime
	id, err := strconv.Atoi(nTaskid)
	if err != nil {
		log.Printf("❌ Task ID Conversion Error: %v", err)
		c.JSON(400, gin.H{"error": "Invalid Task ID"})
		return
	}

	P256dh := sub.Keys.P256dh
	Auth := sub.Keys.Auth
	Endpoint := sub.Endpoint

	notify := Notify{
		CreatedDate: createdDate,
		DateTime:    runAt,
		ToDoID:      id,
		UUID:        uuid,
		Endpoint:    Endpoint,
		P256dh:      P256dh,
		Auth:        Auth,
	}

	if err := db.Model(&Notify{}).Create(&notify).Error; err != nil {
		fmt.Println("Error creating Notify:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notifyの作成に失敗しました"})
		return
	}
	fmt.Println("Sccuess creating Notify")
	push(notify, c, vapidPrivateKey, vapidPublicKey)
	//push(runAt, c, uuid, id, sub.Endpoint, P256dh, Auth, vapidPrivateKey, vapidPublicKey)

	// 例②：今から10秒後
	// runAt := time.Now().Add(10 * time.Second)

	c.JSON(200, gin.H{"success": true, "status": 200})
}
