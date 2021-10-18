package model

import "time"

type BookReader struct {
	ID         uint
	BookID     uint
	ReaderID   uint
	TakenAt    time.Time
	ReturnedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
