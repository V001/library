package model

import "time"

type BookReader struct {
	BookID     string
	ReaderID   string
	TakenAt    time.Time
	ReturnedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
