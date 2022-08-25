package handlers

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type handlers struct {
	db  *gorm.DB
	log *log.Logger
}

func Init(database *gorm.DB, logger *log.Logger) *handlers {
	return &handlers{database, logger}
}

func (h handlers) HandlePing(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal("pong")
	if err != nil {
		h.log.Printf("error while json marshal:%s", err.Error())
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		h.log.Printf("error while writing response:%s", err.Error())
		return
	}
	return
}
