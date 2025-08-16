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

func EditSubject(c *gin.Context) {
	println("subject/edit")
	uuid := GetProfile(c).UUID
	var req struct {
		ID        int    `json:"id"`
		AfterName string `json:"aftername"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	println(req.ID)
	println(req.AfterName)

	res := db.Model(&Subjects{}).Where("id = ? AND uuid = ?", req.ID, uuid).Update("name", req.AfterName)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}

	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "指定された科目が見つかりません"})
		return
	}
	if req.AfterName == "" {
		res = db.Model(&Subjects{}).Where("id = ? AND uuid = ?", req.ID, uuid).Delete(&Subjects{})
	}

	println("Sccuess editing subject")
	c.JSON(http.StatusOK, gin.H{"message": "科目を編集しました"})
}
