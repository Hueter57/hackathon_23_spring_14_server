package model

import "time"

type User struct {
	ID        string    `json:"id" gorm:"size:32;primary_key"`
	DatePoint int       `json:"datePoint"`
	Date      time.Time `json:"date"`
}
