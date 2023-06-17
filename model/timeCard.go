package model

import "time"

type TimeCards struct {
	Date    time.Time   `json:"date"`
	ID      string `json:"id" gorm:"size:32;primary_key"`
	ItemID  string `json:"itemID"`
}