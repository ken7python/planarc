package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TODOLIST struct {
	ID        uint   `gorm:"primaryKey"`
	Date      string `gorm:"not null"`
	SubjectID int    `gorm:"not null"`
	UUID      string `gorm:"not null"`
	Title     string `gorm:"not null"`
	Status    string `gorm:"not null"`
}

func getTODOByUserID(c *gin.Context) {
	fmt.Println("todo/")
	uuid := GetProfile(c).UUID

	date := c.Query("date")

	var todos []TODOLIST

	res := db.Model(&TODOLIST{}).Where("uuid = ? and date = ?", uuid, date).Find(&todos)
	if res.Error != nil {
		fmt.Println("Error fetching ToDo List:", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched ToDo:", len(todos))
	c.JSON(http.StatusOK, todos)
}

func AddToDo(c *gin.Context) {
	fmt.Println("todo/add")

	//body, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(string(body))
	//
	uuid := GetProfile(c).UUID
	var req struct {
		Date      string `json:"date"`
		Title     string `json:"title"`
		Status    string `json:"status"`
		SubjectID int    `json:"subjectID"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	todo := TODOLIST{
		Date:      req.Date,
		Title:     req.Title,
		Status:    req.Status,
		SubjectID: req.SubjectID,
		UUID:      uuid,
	}

	if err := db.Model(&TODOLIST{}).Create(&todo).Error; err != nil {
		fmt.Println("Error creating ToDo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの作成に失敗しました"})
		return
	}
	fmt.Println("Sccuess creating ToDO")
	c.JSON(http.StatusOK, gin.H{"message": "ToDoを作成しました"})
}
