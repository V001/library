package model

import "time"

type Author struct {
	ID        string
	Fio       string
	CreatedAt time.Time
}
