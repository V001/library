package model

import "time"

type BookReader struct {
	BookID     uint
	ReaderID   uint
	TakenAt    time.Time
	ReturnedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
