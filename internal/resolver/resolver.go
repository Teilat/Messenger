package resolver

import (
	"gorm.io/gorm"
	"log"
)

type Resolver struct {
	Log *log.Logger
	Db  *gorm.DB
}

func Init(db *gorm.DB, log *log.Logger) *Resolver {
	return &Resolver{
		Log: log,
		Db:  db,
	}
}
