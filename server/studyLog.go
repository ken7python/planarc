package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StudyLog struct {
	ID           uint   `gorm:"primaryKey"`
	Date         string `gorm:"not null"`
	SubjectID    int    `gorm:"not null"`
	UUID         string `gorm:"not null"`
	StartHours   int    `gorm:"not null"`
	StartMinutes int    `gorm:"not null"`
	EndHours     int    `gorm:"not null"`
	EndMinutes   int    `gorm:"not null"`
	StudyTime    int    `gorm:"not null"`
}

func getLogByUserID(c *gin.Context) {
	fmt.Println("studylog/")
	uuid := GetProfile(c).UUID

	date := c.Query("date")

	var logs []StudyLog

	res := db.Model(&StudyLog{}).Where("uuid = ? and date = ?", uuid, date).Find(&logs)
	if res.Error != nil {
		fmt.Println("Error fetching study logs:", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "勉強記録の取得に失敗しました"})
		return
	}
	fmt.Println("Fetched study logs:", len(logs))
	c.JSON(http.StatusOK, logs)
}

func AddLog(c *gin.Context) {
	fmt.Println("studylog/add")

	//body, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(string(body))
	//
	uuid := GetProfile(c).UUID
	var req struct {
		Date         string `json:"date"`
		StartHours   int    `json:"sHours"`
		StartMinutes int    `json:"sMinutes"`
		EndHours     int    `json:"eHours"`
		EndMinutes   int    `json:"eMinutes"`
		SubjectID    int    `json:"subjectID"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	studying := (req.EndHours-req.StartHours)*60 + (req.EndMinutes - req.StartMinutes)
	log := StudyLog{
		Date:         req.Date,
		StartHours:   req.StartHours,
		StartMinutes: req.StartMinutes,
		EndHours:     req.EndHours,
		EndMinutes:   req.EndMinutes,
		SubjectID:    req.SubjectID,
		StudyTime:    studying,
		UUID:         uuid,
	}
	if studying < 0 {
		fmt.Println("studying < 0")
		c.JSON(http.StatusBadRequest, gin.H{"error": "終了時間は開始時間より後でなければなりません"})
		return
	}
	if err := db.Model(&StudyLog{}).Create(&log).Error; err != nil {
		fmt.Println("Error creating subject:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "勉強記録の作成に失敗しました"})
		return
	}
	fmt.Println("Sccuess creating study log")
	c.JSON(http.StatusOK, gin.H{"message": "勉強記録を作成しました"})
}
