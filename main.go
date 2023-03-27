package main

import (
	"Messenger/internal/cache"
	"Messenger/internal/config"
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"Messenger/internal/resolver"
	"Messenger/webapi"
)

func main() {
	// getting config from config file
	config.GetConf()

	// Initiate and start database connection
	db, err := database.NewDbProvider(logger.NewLogger("[Database] ")).Start()
	// Initiate and start cache
	cache, updChan, delChan := cache.NewCache(logger.NewLogger("[Database]"), db.Database).Start()
	// Start database updater with channels from cache
	db.StartUpdateListener(updChan, delChan)

	// Initiate internal resolvers
	res := resolver.Init(logger.NewLogger("[Resolver]"), cache)

	api := webapi.NewWebapi(logger.NewLogger("[Handler]"), res)
	err = api.Run()
	if err != nil {
		return
	}
}
