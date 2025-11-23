package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type unfinishedLIST struct {
	ID        uint      `gorm:"primaryKey"`
	Date      string    `gorm:"not null"`
	SubjectID int       `gorm:"not null"`
	UUID      string    `gorm:"not null"`
	Title     string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func retGetUnfinishedByUserID(uuid string) []unfinishedLIST {
	var unfinishedLists []unfinishedLIST

	res := db.Model(&unfinishedLIST{}).Where("uuid = ?", uuid).Find(&unfinishedLists)
	if res.Error != nil {
		fmt.Println("Error fetching Unfinished List:", res.Error)
		return nil
	}
	return unfinishedLists
}

func retGetUnfinishedByUserIDTop10(uuid string) []unfinishedLIST {
	var unfinishedLists []unfinishedLIST

	res := db.Model(&unfinishedLIST{}).Where("uuid = ?", uuid).Order("date ASC").Limit(10).Find(&unfinishedLists)
	if res.Error != nil {
		fmt.Println("Error fetching Unfinished List:", res.Error)
		return nil
	}
	return unfinishedLists
}

func getUnfinishedByUserID(c *gin.Context) {
	fmt.Println("unfinished/")
	uuid := GetProfile(c).UUID

	unfinishedLists := retGetUnfinishedByUserID(uuid)
	if unfinishedLists == nil {
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

func deleteUnfinished(c *gin.Context) {
	fmt.Println("unfinished/delete")
	var req struct {
		ID uint `json:"id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	db.Model(&unfinishedLIST{}).Where("id = ?", req.ID).Delete(&unfinishedLIST{})

	fmt.Println("Sccuess deleting Unfinished List")
	c.JSON(http.StatusOK, gin.H{"message": "Unfinished Listを削除しました"})
}

func backUnfinished(c *gin.Context) {
	fmt.Println("unfinished/back")
	var req struct {
		ID   uint   `json:"id"`
		Date string `json:"date"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	var unfinished unfinishedLIST

	res := db.Model(&unfinishedLIST{}).Where("id = ?", req.ID).First(&unfinished)
	if res.Error != nil {
		fmt.Println("Error fetching Unfinished List:", res.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UnfinishedLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched Unfinished List:", unfinished)

	//loc, _ := time.LoadLocation("Asia/Tokyo")
	//now := time.Now().In(loc)

	var todo = TODOLIST{
		//Date:      now.Format("2006-01-02"),
		Date:      req.Date,
		Title:     unfinished.Title,
		Status:    unfinished.Status,
		SubjectID: unfinished.SubjectID,
		UUID:      unfinished.UUID,
		Checked:   false,
	}

	if err := db.Model(&TODOLIST{}).Create(&todo).Error; err != nil {
		fmt.Println("Error creating ToDo List:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDo Listの作成に失敗しました"})
		return
	}

	db.Model(&unfinishedLIST{}).Where("id = ?", req.ID).Delete(&unfinishedLIST{})

	fmt.Println("Sccuess creating ToDo List")
	c.JSON(http.StatusOK, gin.H{"message": "ToDo Listを作成しました"})
}
