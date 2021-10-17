package model

type Publisher struct {
	ID    uint `gorm:"primaryKey"`
	Title string
}
