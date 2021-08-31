package models

import (
	"time"
)

type Author struct {
	ID        uint `gorm:"primaryKey"`
	FIO       string
	CreatedAt time.Time
}
