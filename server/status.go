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
	mood      uint   `gorm:"not null"`
}

func statusInit(c *gin.Context) *Status {
	fmt.Println("todo/")
	uuid := GetProfile(c).UUID

	date := c.Query("date")

	var statusToday Status
	db.Model(&statusToday).Where("UUID = ? AND Date = ? ", uuid, date).Find(&statusToday)
	if len(statusToday.Date) > 0 {
		return &statusToday
	} else {
		statusToday = Status{
			UUID:      uuid,
			Date:      date,
			Enjoyment: "",
			mood:      0,
		}
		if err := db.Model(&statusToday).Create(&statusToday); err != nil {
			fmt.Println("Error creating status:", err)
			return nil
		} else {
			return &statusToday
		}
	}
}
