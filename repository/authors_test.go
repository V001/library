package repository

import (
	"github.com/v001/library/configs"
	db2 "github.com/v001/library/db"
	"github.com/v001/library/model"
	"testing"
)

func TestAuthorRepository_Create(t *testing.T) {
	cfg := configs.Get()
	var err error
	dbCon, err := db2.CreateConnection(cfg)
	if err != nil {
		t.Error(err)
	}
	rep := NewAuthorRepository(dbCon)
	var author model.Author
	var id uint
	if id, err = rep.Create(author); err != nil {
		t.Error(err)
	}
	t.Log(id)

}
