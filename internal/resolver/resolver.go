package resolver

import (
	"Messenger/internal/cache"
	"gorm.io/gorm"
	"log"
)

type Resolver struct {
	Log          *log.Logger
	Db           *gorm.DB
	MessageCache *cache.Cache
}

func Init(db *gorm.DB, log *log.Logger) *Resolver {
	return &Resolver{
		Log: log,
		Db:  db,
	}
}
