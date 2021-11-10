package model

import "time"

type Author struct {
	ID        uint   `gorm:"primaryKey"`
	Fio       string `json:"fio"`
	CreatedAt time.Time
}
