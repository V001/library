package model

type Publisher struct {
	//gorm.Model
	ID    uint   `gorm:"primaryKey" json:"ID"`
	Title string `json:"title"`
}
