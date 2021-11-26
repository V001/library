package storage

import (
	"context"
	"fmt"
	"github.com/v001/library/configs"
	"github.com/v001/library/internal/storage/pg"
	"gorm.io/gorm"
)

type Storage struct {
	Pg *gorm.DB

	Author     IAuthorRepository
	BookAuthor IBookAuthorRepository
	BookReader IBookReaderRepository
	Books      IBookRepository
	Genres     IGenreRepository
	Publisher  IPublisherRepository
	ReaderHall IReaderHallRepository
	Readers    IReaderRepository
	Reports    IReportRepository
}

func New(ctx context.Context, cfg *configs.Config) (*Storage, error) {

	// connect to Postgres
	pgDB, err := pg.Dial(cfg)
	if err != nil {
		return nil, fmt.Errorf("pgdb.Dial failed %v", err)
	}

	// Run Postgres migrations
	//if pgDB != nil {
	//	log.Println("Running PostgreSQL migrations...")
	//	if err := runPgMigrations(); err != nil {
	//		return nil, errors.Wrap(err, "runPgMigrations failed")
	//	}
	//}

	var storage Storage

	// Init Postgres repositories
	if pgDB != nil {
		storage.Pg = pgDB
		storage.Author = pg.NewAuthorRepository(pgDB)
		storage.BookAuthor = pg.NewBookAuthorRepository(pgDB)
		storage.BookReader = pg.NewBookReaderRepository(pgDB)
		storage.Books = pg.NewBookRepository(pgDB)
		storage.Genres = pg.NewGenreRepository(pgDB)
		storage.Publisher = pg.NewPublisherRepository(pgDB)
		storage.ReaderHall = pg.NewReaderHallRepository(pgDB)
		storage.Readers = pg.NewReaderRepository(pgDB)
		storage.Reports = pg.NewReportRepo(pgDB)
	}

	return &storage, nil
}
