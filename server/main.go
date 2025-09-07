// go run main.go database.go user.go
package main

import (
	"sync"

	"time"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

func main() {
	InitDB_MySQL()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:4173", "https://planarc.kencode.tech"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

	todo := api.Group("/todo")
	todo.Use(authMiddleware())
	todo.GET("/", getTODOByUserID)
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

	fmt.Println("Starting server")

	r.Run("0.0.0.0:8080")
}
