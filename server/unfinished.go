package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type unfinishedLIST struct {
	ID        uint   `gorm:"primaryKey"`
	Date      string `gorm:"not null"`
	SubjectID int    `gorm:"not null"`
	UUID      string `gorm:"not null"`
	Title     string `gorm:"not null"`
	Status    string `gorm:"not null"`
}

func getUnfinishedByUserID(c *gin.Context) {
	fmt.Println("unfinished/")
	uuid := GetProfile(c).UUID

	var unfinishedLists []unfinishedLIST

	res := db.Model(&unfinishedLIST{}).Where("uuid = ?", uuid).Find(&unfinishedLists)
	if res.Error != nil {
		fmt.Println("Error fetching Unfinished List:", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UnfinishedLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched Unfinished List:", len(unfinishedLists))
	c.JSON(http.StatusOK, unfinishedLists)
}

func moveToUnfinished(c *gin.Context) {
	fmt.Println("unfinished/move")
	var req struct {
		ID uint `json:"id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	var todo TODOLIST

	res := db.Model(&TODOLIST{}).Where("id = ?", req.ID).First(&todo)
	if res.Error != nil {
		fmt.Println("Error fetching ToDo List:", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched ToDo:", todo)

	var unfinished unfinishedLIST = unfinishedLIST{
		Date:      todo.Date,
		Title:     todo.Title,
		Status:    todo.Status,
		SubjectID: todo.SubjectID,
		UUID:      todo.UUID,
	}

	if err := db.Model(&unfinishedLIST{}).Create(&unfinished).Error; err != nil {
		fmt.Println("Error creating Unfinished List:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unfinished Listの作成に失敗しました"})
		return
	}

	db.Model(&TODOLIST{}).Where("id = ?", req.ID).Delete(&TODOLIST{})

	fmt.Println("Sccuess creating Unfinished List")
	c.JSON(http.StatusOK, gin.H{"message": "Unfinished Listを作成しました"})
}
