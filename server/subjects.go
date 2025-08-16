package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Subjects struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique/not null"`
	UUID  string `gorm:"unique/not null"`
	Color string `gorm:"unique/not null"`
}

func getSubjectByUserID(c *gin.Context) {
	println("subject/")
	uuid := GetProfile(c).UUID
	var subjectList []Subjects

	if err := db.Where("uuid = ?", uuid).Find(&subjectList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "科目の取得に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, subjectList)
}

func AddSubject(c *gin.Context) {
	println("subject/add")
	uuid := GetProfile(c).UUID
	var req struct {
		Name  string `json:"name"`
		Color string `json:"color"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
	}
	subject := Subjects{
		Name:  req.Name,
		Color: req.Color,
		UUID:  uuid,
	}
	if err := db.Create(&subject).Error; err != nil {
		fmt.Println("Error creating subject:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サブジェクトの作成に失敗しました"})
	}
	println("Sccuess creating subject")
	c.JSON(http.StatusOK, gin.H{"message": "科目を作成しました"})
}
