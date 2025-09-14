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
	Checked   bool   `gorm:"not null"`
	Status    string `gorm:"not null"`
}

func retGetTODOByUserID(uuid string, date string) []TODOLIST {
	var todos []TODOLIST

	res := db.Model(&TODOLIST{}).Where("uuid = ? and date = ?", uuid, date).Find(&todos)
	if res.Error != nil {
		fmt.Println("Error fetching ToDo List:", res.Error)
		return nil
	}
	return todos
}

func retGetTODO_Checked(uuid string, date string) []TODOLIST {
	var todos []TODOLIST

	res := db.Model(&TODOLIST{}).Where("uuid = ? and date = ? and checked = 1", uuid, date).Find(&todos)
	if res.Error != nil {
		fmt.Println("Error fetching ToDo List:", res.Error)
		return nil
	}
	return todos
}

func retGetTODObyStatusUnfinished(uuid string, date string, status string) []TODOLIST {
	var todos []TODOLIST

	res := db.Model(&TODOLIST{}).Where(`uuid = ? and date = ? and checked = 0 and status = '`+status+`'`, uuid, date).Find(&todos)
	if res.Error != nil {
		fmt.Println("Error fetching ToDo List:", res.Error)
		return nil
	}
	return todos
}

func getTODOByUserID(c *gin.Context) {
	fmt.Println("todo/")
	uuid := GetProfile(c).UUID

	date := c.Query("date")

	todos := retGetTODOByUserID(uuid, date)
	if todos == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}

	fmt.Println("Fetched ToDo:", len(todos))
	c.JSON(http.StatusOK, todos)
}

func getToDOByGroup(c *gin.Context) {
	fmt.Println("todo/group")
	uuid := GetProfile(c).UUID

	date := c.Query("date")

	checked := retGetTODO_Checked(uuid, date)
	if checked == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched ToDo:", len(checked))

	MUST := retGetTODObyStatusUnfinished(uuid, date, "MUST")
	if MUST == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched ToDo:", len(MUST))

	WANT := retGetTODObyStatusUnfinished(uuid, date, "WANT")
	if WANT == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "TODOLISTの取得に失敗しました"})
		return
	}
	fmt.Println("Fetched ToDo:", len(WANT))

	c.JSON(http.StatusOK, gin.H{"checked": checked, "MUST": MUST, "WANT": WANT})
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
		Checked:   false,
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

func ToDoChecked(c *gin.Context) {
	fmt.Println("todo/checked")
	var req struct {
		ID int `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}

	var todo TODOLIST
	if err := db.Model(&TODOLIST{}).Where("id = ?", req.ID).First(&todo).Error; err != nil {
		fmt.Println("Error fetching ToDo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの取得に失敗しました"})
		return
	}

	todo.Checked = !todo.Checked

	if err := db.Save(&todo).Error; err != nil {
		fmt.Println("Error updating ToDo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの更新に失敗しました"})
		return
	}

	fmt.Println("Sccuess updating ToDO")
	c.JSON(http.StatusOK, gin.H{"message": "ToDoを更新しました"})
}

func ToDoEdit(c *gin.Context) {
	fmt.Println("todo/edit")
	var req struct {
		ID       int    `json:"id"`
		NewTitle string `json:"newtitle"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの解析に失敗しました"})
		return
	}
	var todo TODOLIST
	if err := db.Model(&TODOLIST{}).Where("id = ?", req.ID).
		First(&todo).Error; err != nil {
		fmt.Println("Error fetching ToDo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの取得に失敗しました"})
		return
	}
	todo.Title = req.NewTitle

	if todo.Title == "" {
		if err := db.Model(&TODOLIST{}).Where("id = ?", req.ID).Delete(&TODOLIST{}).Error; err != nil {
			fmt.Println("Error deleting ToDo:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの削除に失敗しました"})
			return
		}
		fmt.Println("ToDo deleted successfully")
	} else {
		if err := db.Save(&todo).Error; err != nil {
			fmt.Println("Error updating ToDo:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ToDoの更新に失敗しました"})
			return
		}
		fmt.Println("Sccuess updating ToDO")
	}
	c.JSON(http.StatusOK, gin.H{"message": "ToDoを更新しました"})
}
