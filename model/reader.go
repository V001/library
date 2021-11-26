package model

import "time"

type Reader struct {
	ID       uint      `gorm:"primaryKey" json:"ID,omitempty"`
	Fio      string    `json:"fio,omitempty"`
	Birthday time.Time `json:"birthday"`
	Phone    string    `json:"phone,omitempty"`
}
