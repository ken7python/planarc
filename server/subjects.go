package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Subjects struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique/not null"`
	UUID      string    `gorm:"unique/not null"`
	Color     string    `gorm:"unique/not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func retGetSubjectByUserID(uuid string) []Subjects {
	var subjectList []Subjects

	if err := db.Where("uuid = ?", uuid).Find(&subjectList).Error; err != nil {
		return nil
	}
	return subjectList
}

func getSubjectByUserID(c *gin.Context) {
	fmt.Println("subject/")
	uuid := GetProfile(c).UUID

	subjectList := retGetSubjectByUserID(uuid)
	if subjectList == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "科目の取得に失敗しました"})
	}

	c.JSON(http.StatusOK, subjectList)
}

func AddSubject(c *gin.Context) {
	fmt.Println("subject/add")
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
	fmt.Println("Sccuess creating subject")
	c.JSON(http.StatusOK, gin.H{"message": "科目を作成しました"})
}

func EditSubject(c *gin.Context) {
	fmt.Println("subject/edit")
	uuid := GetProfile(c).UUID
	var req struct {
		ID         int    `json:"id"`
		AfterName  string `json:"aftername"`
		AfterColor string `json:"aftercolor"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	fmt.Println(req.ID)
	fmt.Println(req.AfterName)

	res := db.Model(&Subjects{}).Where("id = ? AND uuid = ?", req.ID, uuid).Update("name", req.AfterName).Update("color", req.AfterColor)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	if req.AfterName == "" {
		res = db.Model(&Subjects{}).Where("id = ? AND uuid = ?", req.ID, uuid).Delete(&Subjects{})
	}

	fmt.Println("Sccuess editing subject")
	c.JSON(http.StatusOK, gin.H{"message": "科目を編集しました"})
}
