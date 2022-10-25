package main

import (
	"Messenger/db"
	"Messenger/internal/config"
	"Messenger/webapi"
)

func main() {
	// getting config from config file
	config.GetConf()
	db, err := db.InitPostgresql()
	if err != nil {
		return
	}
	err = webapi.Run(db)
	if err != nil {
		return
	}
}
