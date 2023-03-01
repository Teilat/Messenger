package resolver

import (
	"Messenger/internal/cache"
	"Messenger/internal/logger"
)

type Resolver struct {
	*logger.Logger
	MessageCache cache.Cache
}

func Init(log *logger.Logger, cache cache.Cache) *Resolver {
	return &Resolver{
		Logger:       log,
		MessageCache: cache,
	}
}
