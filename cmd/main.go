package main

import (
	"github.com/v001/library/configs"
	"github.com/v001/library/db"
	"github.com/v001/library/transport/http"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	conf := configs.Get()

	//init DB
	_, err := db.ConnectToDB(&conf.DB)
	if err != nil {
		return err
	}

	//init server
	srv := http.NewServer()
	if conf.HTTP.Enable {
		if err := srv.Run(conf.HTTP); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}
