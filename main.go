package main

import (
	"Messenger/database"
	"Messenger/internal/config"
	"Messenger/webapi"
)

func main() {
	// getting config from config file
	config.GetConf()
	db, err := database.InitPostgresql()
	if err != nil {
		return
	}
	err = webapi.Init(db)
	if err != nil {
		return
	}
}
