package main

import (
	"Messenger/internal/cache"
	"Messenger/internal/config"
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"Messenger/internal/resolver"
	"Messenger/webapi"
	"Messenger/webapi/handlers"
)

func main() {
	// getting config from config file
	config.GetConf()

	dbProvider := database.NewDbProvider(logger.NewLogger("[Database] "))
	db := dbProvider.Run()

	cache := cache.NewCache(logger.NewLogger("[Database]"), db)
	cache.Start()

	hub := resolver.NewHub()

	resolverLog := logger.NewLogger("[Resolver]")
	res := resolver.Init(resolverLog, cache)

	handlerLog := logger.NewLogger("[Handler]")
	h := handlers.Init(handlerLog, res, hub)

	api := webapi.NewWebapi(db, res, hub, h)
	err := api.Run()
	if err != nil {
		return
	}
}
