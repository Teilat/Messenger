package main

import (
	"Messenger/database"
	"Messenger/internal/cache"
	"Messenger/internal/config"
	"Messenger/internal/logger"
	"Messenger/internal/resolver"
	"Messenger/webapi"
	"Messenger/webapi/handlers"
	"log"
	"os"
)

func main() {
	// getting config from config file
	config.GetConf()

	dbProvider := database.NewDbProvider(logger.NewLogger(log.New(os.Stderr, "[Database] ", log.LstdFlags)))
	db := dbProvider.Run()

	cacheProvider := cache.NewCache(logger.NewLogger(log.New(os.Stderr, "[Database] ", log.LstdFlags)), db)
	cacheProvider.Run()

	hub := resolver.NewHub()

	resolverLog := logger.NewLogger(log.New(os.Stderr, "[Resolver] ", log.LstdFlags))
	res := resolver.Init(db, resolverLog)

	handlerLog := logger.NewLogger(log.New(os.Stderr, "[Handler] ", log.LstdFlags))
	h := handlers.Init(handlerLog, res, hub)

	api := webapi.NewWebapi(db, res, hub, h)
	err := api.Run()
	if err != nil {
		return
	}
}
