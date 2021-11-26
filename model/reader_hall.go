package model

type ReaderHall struct {
	ID    uint   `gorm:"primaryKey" json:"ID,omitempty"`
	Title string `json:"title,omitempty"`
}
