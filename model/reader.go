package model

import "time"

type Reader struct {
	ID       string
	Fio      string
	Birthday time.Time
	Phone    string
}
