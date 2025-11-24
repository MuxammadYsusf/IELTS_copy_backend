package main

import (
	"fmt"
	"log"
	"net/http"

	"github/http/copy/test_task1/config"
	"github/http/copy/test_task1/controller"
	"github/http/copy/test_task1/pkg/db"
)

func main() {
	cfg := config.Load()

	// simple log
	conn, err := db.NewConnectPostgres(&cfg)
	log.Println("connection:", conn)
	if err != nil {
		fmt.Println(err, "this is the error")
		log.Fatal(err)
	}

	cont := controller.NewController(conn)

	http.HandleFunc("/movie/", cont.Movie)

	log.Println("Listening...", cfg.HTTPPort)

	err = http.ListenAndServe(cfg.HTTPPort, nil)
	if err != nil {
		log.Fatal("ERROR!", err)
		return
	}
}
