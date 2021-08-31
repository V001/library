package models

import "time"

type Book struct {
	ID             uint `gorm:"primaryKey"`
	Title          string
	PublishedTime  time.Time
	PublishedPlace string
	PageCnt        uint
	HallID         string
	CreatedAt      time.Time
	GenreID        string
	PublisherID    string
}
