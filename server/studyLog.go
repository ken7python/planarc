package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudyLog struct {
	ID        uint   `gorm:"primaryKey"`
	subjectID string `gorm:"unique/not null"`
	UUID      string `gorm:"unique/not null"`
	sHours    int    `gorm:"unique/not null"`
	sMinutes  int    `gorm:"unique/not null"`
	eHours    int    `gorm:"unique/not null"`
	eMinutes  int    `gorm:"unique/not null"`
	studyTime int    `gorm:"unique/not null"`
}

func getLogByUserID(c *gin.Context) {
	println("subject/")
	uuid := GetProfile(c).UUID
	var Logs []StudyLog

	if err := db.Where("uuid = ?", uuid).Find(&Logs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "勉強記録の取得に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, Logs)
}

func AddLog(c *gin.Context) {
	println("subject/add")
	uuid := GetProfile(c).UUID
	var req struct {
		sHours   int `json:"sHours"`
		sMinutes int `json:"sMintes"`
		eHours   int `json:"eHours"`
		eMinutes int `json:"eMinutes"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
	}

	studying := (req.eHours-req.sHours)*60 + (req.eMinutes - req.sMinutes)
	log := StudyLog{
		sHours:    req.sHours,
		sMinutes:  req.sMinutes,
		eHours:    req.eHours,
		eMinutes:  req.eMinutes,
		studyTime: studying,
		UUID:      uuid,
	}
	if studying < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "終了時間は開始時間より後でなければなりません"})
		return
	}
	if err := db.Create(&log).Error; err != nil {
		fmt.Println("Error creating subject:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "勉強記録の作成に失敗しました"})
	}
	println("Sccuess creating subject")
	c.JSON(http.StatusOK, gin.H{"message": "勉強記録を作成しました"})
}
