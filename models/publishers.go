package models

type Publisher struct {
	ID    uint `gorm:"primaryKey"`
	Title string
}
