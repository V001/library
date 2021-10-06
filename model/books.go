package model

import "time"

type Book struct {
	ID             string
	Title          string
	PublishedTime  time.Time
	PublisherPlace string
	PageCnt        int32
	HallID         string
	CreatedAt      time.Time
	genreID        string
	publisherID    string
}
