package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func retGetAnalysis(uuid string, startStr string, endStr string) []StudyLog {
	loc, _ := time.LoadLocation("Asia/Tokyo")

	start := dateFormat(startStr).In(loc)

	fmt.Printf("Start Date: %v\n", start)

	end := dateFormat(endStr).In(loc)
	fmt.Printf("End Date: %v\n", end)

	logs := retGetLogByUserIDAllDay(uuid)
	logsInRange := []StudyLog{}
	for _, log := range logs {
		//fmt.Println(log.Date)
		logDate := dateFormat(log.Date).In(loc)
		//fmt.Println(logDate)
		if inRange(logDate, start, end) {
			logsInRange = append(logsInRange, log)
		}
	}
	//fmt.Printf("Logs in Range: %v\n", len(logsInRange))
	return logsInRange
}

func getAnalysis(c *gin.Context) {
	//startStr := "2025-01-01"
	//endStr := "2025-10-01"
	startStr := c.Query("start")
	endStr := c.Query("end")
	fmt.Println(startStr, endStr)
	if len(startStr) == 0 || len(endStr) == 0 {
		c.JSON(400, gin.H{"error": "startとendの両方のパラメータが必要です"})
		return
	}

	uuid := GetProfile(c).UUID
	if len(uuid) == 0 {
		c.JSON(400, gin.H{"error": "UUIDの取得に失敗しました"})
	}

	logs := retGetAnalysis(uuid, startStr, endStr)
	fmt.Println(logs)

	c.JSON(200, logs)
}

func dateFormat(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(err)
	}
	return t
}

func inRange(date, start, end time.Time) bool {
	return (date.Equal(start) || date.After(start)) && (date.Equal(end) || date.Before(end))
}
