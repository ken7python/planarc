package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY is not set")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendGeminiAPI, // AI StudioのGemini APIを使う
		APIKey:  apiKey,
	})
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Close()

	contents := []*genai.Content{{
		Parts: []*genai.Part{{Text: "あなたは私の学校の先生です。私が勉強をがんばれるような声をかけるのが得意です。今日の勉強時間は2時間30分です。完了したのは英語の長文問題集とインターンシップの作文です。未完了は数学IAのワークと英単語です。振り返り文には課題がたまっていたけれど、少し進められてよかったと書いていました。手帳のコメントとしてメッセージをお願いします"}},
	}}

	resp, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", contents, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, cand := range resp.Candidates {
		for _, p := range cand.Content.Parts {
			if p.Text != "" {
				fmt.Println(p.Text)
			}
		}
	}
}
