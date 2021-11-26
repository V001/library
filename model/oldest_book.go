package model

type OldestBook struct {
	Title     string `json:"book_title" gorm:"column:book_title" `
	Age       int    `json:"age"`
	HallTitle string `json:"hall_title" gorm:"column:hall_title"`
}
