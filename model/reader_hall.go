package model

type ReaderHall struct {
	ID    uint `gorm:"primaryKey"`
	Title string
}
