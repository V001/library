package model

type PopulatedBooks struct {
	ID       int    `json:"book_id" gorm:"column:book_id"`
	Title    string `gorm:"column:book_title" json:"book_title"`
	TakenCnt int    `json:"taken_cnt"`
}
