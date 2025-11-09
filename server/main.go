// go run main.go database.go user.go
package main

import (
	"github.com/SherClockHolmes/webpush-go"
	"io"
	"log"
	"sync"

	"fmt"

	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

var (
	limiter = rate.NewLimiter(4, 16) // 3 requests per second with a burst of 10
	mu      sync.Mutex
)

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		if !limiter.Allow() {
			c.JSON(429, gin.H{"error": "たくさんのリクエストを送信しています。しばらくしてから再試行してください。"})
			c.Abort()
			return
		}
		c.Next()
	}
}

type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
}

func main() {
	InitDB_MySQL()
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	CORS_GO := os.Getenv("CORS_GO")

	// ローカルのときは有効にしてください
	if CORS_GO == "ON" {
		r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4173", "https://planarc.kencode.tech", "https://planarc.kencode.tech/"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
			AllowOriginFunc: func(origin string) bool {
				return strings.Contains(origin, "planarc.kencode.tech")
			},
		}))
		r.OPTIONS("/*path", func(c *gin.Context) {
			c.Status(204)
		})
	}

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		//c.File("templates/index.html")
		c.JSON(200, gin.H{
			"message": "Welcome to the PlanArc API",
		})
	})
	//r.GET("/signup", func(c *gin.Context) {
	//	c.File("templates/signup.html")
	//})
	//r.GET("/login", func(c *gin.Context) {
	//	c.File("templates/login.html")
	//})

	api := r.Group("/api")
	//api.Use(rateLimitMiddleware())

	accounts := api.Group("/accounts")
	accounts.POST("/register", register)
	accounts.POST("/login", login)
	accounts.GET("/profile", authMiddleware(), profile)
	accounts.Use(rateLimitMiddleware())

	subjects := api.Group("/subject")
	subjects.Use(authMiddleware())
	subjects.GET("/", getSubjectByUserID)
	subjects.POST("/add", AddSubject)
	subjects.POST("/edit", EditSubject)

	studyLogs := api.Group("/studylog")
	studyLogs.Use(authMiddleware())
	studyLogs.GET("/", getLogByUserID)
	studyLogs.POST("/add", AddLog)
	studyLogs.POST("/delete", deleteLogByID)

	todo := api.Group("/todo")
	todo.Use(authMiddleware())
	todo.GET("/", getTODOByUserID)
	todo.GET("/group", getToDOByGroup)
	todo.POST("/add", AddToDo)
	todo.POST("/check", ToDoChecked)
	todo.POST("/edit", ToDoEdit)

	unfinished := api.Group("/unfinished")
	unfinished.Use(authMiddleware())
	unfinished.GET("/", getUnfinishedByUserID)
	unfinished.POST("/move", moveToUnfinished)
	unfinished.POST("/delete", deleteUnfinished)
	unfinished.POST("/back", backUnfinished)

	status := api.Group("/status")
	status.Use(authMiddleware())
	status.GET("/", getStatus)
	status.POST("/enjoyment", setEnjoyment)
	status.POST("/mood", setMood)

	gemini := api.Group("/comment")
	gemini.Use(authMiddleware())
	gemini.POST("/ask", reqComment)
	gemini.GET("/", getComment)

	analysys := api.Group("/analysis")
	analysys.Use(authMiddleware())
	analysys.GET("/", getAnalysis)

	api.POST("/send", func(c *gin.Context) {
		var sub Subscription
		body, _ := io.ReadAll(c.Request.Body)
		json.Unmarshal(body, &sub)

		// 通知内容
		message := map[string]string{
			"title": "🎉 GoからWeb Push通知！",
			"body":  "こんにちは！Goサーバーから届いたよ！",
		}
		payload, _ := json.Marshal(message)

		// WebPush送信
		resp, err := webpush.SendNotification(payload, &webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				P256dh: sub.Keys.P256dh,
				Auth:   sub.Keys.Auth,
			},
		}, &webpush.Options{
			VAPIDPrivateKey: os.Getenv("VAPID_PRIVATE_KEY"),
			VAPIDPublicKey:  os.Getenv("VAPID_PUBLIC_KEY"),
			TTL:             30,
			Subscriber:      "Subscriber: \"mailto:test@example.com\"",
		})
		if err != nil {
			log.Println("Error:", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		c.JSON(200, gin.H{"success": true})
	})

	fmt.Println("Starting server")

	r.Run("0.0.0.0:8080")
}
