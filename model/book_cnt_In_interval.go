package model

type BookCntInInterval struct {
	Period string `json:"period" gorm:"column:period"`
	Count  int    `json:"count"`
}
