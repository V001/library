package model

import "time"

type CreateBookReader struct {
	BookID    int `json:"book_id" gorm:"column:book_id"`
	ReaderID  int `json:"reader_id" gorm:"column:reader_id"`
	TakenAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookReader struct {
	ID         uint
	BookID     int    `json:"book_id" gorm:"column:book_id"`
	ReaderID   int    `json:"reader_id" gorm:"column:reader_id"`
	BookTitle  string `json:"book_title" gorm:"column:book_title"`
	TakenAt    time.Time
	ReturnedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
