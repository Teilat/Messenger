package webapi

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var instance *handler

type handler struct {
	db     *gorm.DB
	router *mux.Router
}

func Init(db *gorm.DB) *handler {
	once.Do(func() {

		if instance == nil {
			instance = &handler{
				db:     db,
				router: mux.NewRouter(),
			}
		}
	})
	return instance
}
