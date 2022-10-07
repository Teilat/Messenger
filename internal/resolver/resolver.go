package resolver

import (
	"Messenger/internal/cache"
	"Messenger/internal/logger"
	"gorm.io/gorm"
)

type Resolver struct {
	Log          *logger.Log
	Db           *gorm.DB
	MessageCache *cache.Cache
}

func Init(db *gorm.DB, log *logger.Log) *Resolver {
	return &Resolver{
		Log: log,
		Db:  db,
	}
}
