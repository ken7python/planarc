// go run main.go database.go user.go
package main

import (
	"sync"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	limiter = rate.NewLimiter(1, 5) // Allow 1 request per second with a burst of 5
	mu      sync.Mutex
)

func rateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		defer mu.Unlock()
		if !limiter.Allow() {
			c.JSON(429, gin.H{"error": "Too many requests, please try again later."})
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
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.File("templates/index.html")
	})
	r.GET("/signup", func(c *gin.Context) {
		c.File("templates/signup.html")
	})
	r.GET("/login", func(c *gin.Context) {
		c.File("templates/login.html")
	})

	api := r.Group("/api")
	api.Use(rateLimitMiddleware())

	accounts := api.Group("/accounts")
	accounts.POST("/register", register)
	accounts.POST("/login", login)
	accounts.GET("/profile", authMiddleware(), profile)

	r.Run("0.0.0.0:8080")
}
