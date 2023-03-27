package resolver

import (
	"Messenger/internal/cache"
	"Messenger/internal/logger"
)

type Resolver struct {
	Logger *logger.Logger
	Cache  cache.Cache
}

func Init(log *logger.Logger, cache cache.Cache) *Resolver {
	return &Resolver{
		Logger: log,
		Cache:  cache,
	}
}
