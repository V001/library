package model

type Genre struct {
	ID    uint `gorm:"primaryKey"`
	Title string
}
