package models

type Genres struct {
	ID    uint `gorm:"primaryKey"`
	Title string
}
