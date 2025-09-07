package main

type Status struct {
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"not null"`
	Date      string `gorm:"not null"`
	Enjoyment string `gorm:"not null"`
	mood      uint   `gorm:"not null"`
}
