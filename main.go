package main

import (
	"Messenger/internal/cache"
	"Messenger/internal/config"
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"Messenger/internal/resolver"
	"Messenger/webapi"
	"context"
)

func main() {
	ctx := context.Background()
	// getting config from config file
	config.GetConf()

	// Initiate and start database connection
	db, err := database.NewDbProvider(logger.NewLogger("[Database] ")).Start()
	if err != nil {
		panic(err)
	}
	// Initiate and start cache
	cache, updChan, delChan := cache.NewCache(logger.NewLogger("[Cache]"), db.GetSnapshot).Start()
	// Start database updater with channels from cache
	db.StartUpdateListener(ctx, updChan, delChan)

	// Initiate internal resolvers
	res := resolver.Init(logger.NewLogger("[Resolver]"), cache)

	api := webapi.NewWebapi(logger.NewLogger("[Handler]"), res)
	err = api.Run()
	if err != nil {
		panic(err)
	}
}
