package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Status struct {
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"not null"`
	Date      string `gorm:"not null"`
	Enjoyment string `gorm:"not null"`
	Mood      uint   `gorm:"not null"`
}

func statusInit(uuid string, date string) *Status {

	var statusToday Status
	db.Model(&statusToday).Where("UUID = ? AND Date = ? ", uuid, date).Find(&statusToday)
	if len(statusToday.Date) > 0 {
		return &statusToday
	} else {
		statusToday = Status{
			UUID:      uuid,
			Date:      date,
			Enjoyment: "",
			Mood:      0,
		}
		if err := db.Model(&Status{}).Create(&statusToday).Error; err != nil {
			fmt.Println("Error creating status:", err)
			return nil
		} else {
			return &statusToday
		}
	}
}

func setEnjoyment(c *gin.Context) {
	uuid := GetProfile(c).UUID
	date := c.Query("date")

	fmt.Println("status/enjoyment")
	statusToday := statusInit(uuid, date)
	if statusToday == nil {
		c.JSON(500, gin.H{"error": "ステータスの初期化に失敗しました"})
		return
	}
	var req struct {
		Enjoyment string `json:"enjoyment"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "リクエストの解析に失敗しました"})
		fmt.Println(err)
		return
	}
	statusToday.Enjoyment = req.Enjoyment
	fmt.Println(req.Enjoyment)
	if err := db.Model(&statusToday).Updates(&statusToday).Error; err != nil {
		c.JSON(500, gin.H{"error": "ステータスの更新に失敗しました"})
		return
	}
	c.JSON(200, gin.H{"message": "ステータスを更新しました"})
}

func setMood(c *gin.Context) {
	fmt.Println("status/mood")

	uuid := GetProfile(c).UUID
	date := c.Query("date")
	statusToday := statusInit(uuid, date)
	if statusToday == nil {
		c.JSON(500, gin.H{"error": "ステータスの初期化に失敗しました"})
		return
	}
	var req struct {
		Mood uint `json:"mood"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "リクエストの解析に失敗しました"})
		fmt.Println(err)
		return
	}
	statusToday.Mood = req.Mood
	fmt.Println(req.Mood)
	if err := db.Model(&statusToday).Updates(&statusToday).Error; err != nil {
		c.JSON(500, gin.H{"error": "ステータスの更新に失敗しました"})
		return
	}
	c.JSON(200, gin.H{"message": "ステータスを更新しました"})
}

func retGetStatus(uuid string, date string) *Status {
	fmt.Println("status")

	statusToday := statusInit(uuid, date)
	if statusToday == nil {
		return nil
	}
	return statusToday
}

func getStatus(c *gin.Context) {
	uuid := GetProfile(c).UUID
	date := c.Query("date")

	res := retGetStatus(uuid, date)
	if res == nil {
		c.JSON(500, gin.H{"error": "ステータスの初期化に失敗しました"})
		return
	}

	c.JSON(200, res)
}
