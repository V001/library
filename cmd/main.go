package main

import (
	"github.com/v001/library/conf"
	"github.com/v001/library/transport/http"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	conf := conf.Get()

	//init server
	srv := http.NewServer()
	if conf.HTTP.Enable {
		if err := srv.Run(conf.HTTP); err != nil {
			log.Fatalln(err)
		}
	}
	return nil
}
