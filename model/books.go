package model

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type CreateBookReq struct {
	Book
}

type Book struct {
	ID             uint      `gorm:"primaryKey" json:"id,omitempty"`
	Title          string    `json:"title"`
	PublishedTime  TimeDate  `json:"published_time" gorm:"published_time"`
	PublishedPlace string    `json:"published_place,omitempty"`
	PageCnt        int32     `json:"page_cnt,omitempty"`
	HallID         string    `json:"hall_id,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	GenreID        string    `json:"genre_id,omitempty"`
	PublisherID    string    `json:"publisher_id,omitempty"`
	Authors        []string  `json:"authors"`
}

type BookListResponse struct {
	ID             uint     `gorm:"column:book_id;" json:"book_id"`
	BookTitle      string   `json:"book_title,omitempty"`
	PublishedTime  TimeDate `json:"published_time" gorm:"published_time"`
	PublishedPlace string   `json:"published_place,omitempty"`
	PageCnt        int32    `json:"page_cnt,omitempty"`
	HallTitle      string   `json:"hall_title,omitempty"`
	GenreTitle     string   `json:"genre_title,omitempty"`
	PublisherTitle string   `json:"publisher_title" gorm:"publisher_title"`
	AuthorsStr     string   `json:"authors" gorm:"column:authors"`
	Readers        string   `json:"readers" gorm:"column:readers"`
	BookReadersID  string   `json:"book_reader_id" gorm:"column:book_reader_id"`
}

type TimeDate struct {
	time.Time
}

func (t *TimeDate) UnmarshalJSON(b []byte) (err error) {
	layout := "2006-01-02"

	s := strings.Trim(string(b), "\"") // remove quotes
	if s == "null" {
		return
	}
	t.Time, err = time.Parse(layout, s)
	return
}

// MarshalJSON marshals the enum as a quoted json string
func (t TimeDate) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(t.Format("2006-01-02"))
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (t *TimeDate) Scan(value interface{}) error {
	result, ok := value.(time.Time)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	*t = TimeDate{result}
	return nil
}

// Value return json value, implement driver.Valuer interface
func (t TimeDate) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	timeStr := t.Format("2006-01-02")
	_ = timeStr
	return timeStr, nil
}
