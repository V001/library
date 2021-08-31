package models

import "time"

type BookReaders struct {
	ID         uint `gorm:"primaryKey"`
	BookID     string
	ReaderID   string
	TakenAt    time.Time
	ReturnedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
