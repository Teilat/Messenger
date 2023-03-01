package resolver

import (
	"Messenger/internal/cache"
	"Messenger/internal/logger"
	"gorm.io/gorm"
)

type Resolver struct {
	*logger.Logger
	Db           *gorm.DB
	MessageCache *cache.Cache
}

func Init(db *gorm.DB, log *logger.Logger, cache *cache.Cache) *Resolver {
	return &Resolver{
		Logger:       log,
		Db:           db,
		MessageCache: cache,
	}
}
