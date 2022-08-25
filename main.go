package main

import (
	"Messenger/database"
	"Messenger/internal/config"
	"Messenger/webapi"
	"fmt"
)

func main() {
	// getting config from config file
	config.GetConf()
	db, err := database.InitPostgresql()
	if err != nil {
		return
	}
	webapi.Init(db)

	fmt.Println(db.Migrator().GetTables())
}
