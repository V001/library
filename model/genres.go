package model

type Genre struct {
	ID    uint   `gorm:"primaryKey" json:"ID,omitempty"`
	Title string `json:"title,omitempty"`
}
