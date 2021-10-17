package model

import "time"

type Reader struct {
	ID       uint `gorm:"primaryKey"`
	Fio      string
	Birthday time.Time
	Phone    string
}
