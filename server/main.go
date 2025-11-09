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

// 指定時刻に非同期で1回だけ実行（過ぎてたら何もしない）
func ScheduleOnce(when time.Time, job func()) {
	d := time.Until(when)
	if d <= 0 {
		fmt.Println("⏱️ もう過ぎてるので実行しません:", when)
		return
	}

	// ゴルーチンで裏実行
	go func() {
		time.Sleep(d)
		job()
	}()
}

func main() {
	InitDB_MySQL()
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}
	CORS_GO := os.Getenv("CORS_GO")

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

		loc, _ := time.LoadLocation("Asia/Tokyo")

		// 例①：特定日時で
		runAt := time.Now().Add(10 * time.Second)

		// 例②：今から10秒後
		// runAt := time.Now().Add(10 * time.Second)

		ScheduleOnce(runAt, func() {
			fmt.Println("🟢 実行しました！:", time.Now().In(loc))

			// 通知内容
			message := map[string]string{
				"title": "🎉 GoからWeb Push通知！",
				"body":  "こんにちは！Goサーバーから届いたよ！",
			}
			payload, _ := json.Marshal(message)
			log.Printf("📝 Payload: %s", string(payload))

			// WebPush送信
			log.Println("🚀 Sending notification...")
			resp, err := webpush.SendNotification(payload, &webpush.Subscription{
				Endpoint: sub.Endpoint,
				Keys: webpush.Keys{
					P256dh: sub.Keys.P256dh,
					Auth:   sub.Keys.Auth,
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
		})

		fmt.Println("✅ スケジュール登録:", runAt)

		c.JSON(200, gin.H{"success": true, "status": 200})
	})

	fmt.Println("Starting server")

	r.Run("0.0.0.0:8080")
}
