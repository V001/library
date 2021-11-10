package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/v001/library/configs"
	"github.com/v001/library/internal/storage"
	"github.com/v001/library/internal/transport/handlers"
	"github.com/v001/library/internal/transport/http"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	conf := configs.Get()
	ctx, cancel := context.WithCancel(context.Background())
	repo, err := storage.New(ctx, conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	l := logrus.New()
	defer cancel()
	if conf.HTTP.Enable {
		hManager := handlers.New(l, repo)
		//init server
		srv := http.NewServer(hManager, conf)
		if err := srv.Run(); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}
