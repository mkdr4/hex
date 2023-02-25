package main

import (
	"github.com/mkdr4/hex/api/router"
	"github.com/mkdr4/hex/api/server"
	"github.com/mkdr4/hex/config"
	"github.com/mkdr4/hex/pkg/database"
	"github.com/mkdr4/hex/pkg/log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatal("config:", err.Error())
	}

	logger, err := log.Setup()
	if err != nil {
		log.Fatal("logger", err.Error())
	}
	defer logger.Close()

	db, err := database.Init()
	if err != nil {
		log.Fatal(log.DB, err.Error())
	}
	defer db.Close()

	srv := server.Init()

	router.Setup(srv, db)

	server.Run(srv)
}
