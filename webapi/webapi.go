package webapi

import (
	"Messenger/webapi/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func Init(database *gorm.DB) error {
	// swag init --parseDependency --parseInternal -g webapi.go
	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))
	logger := log.New(os.Stderr, "SQL: ", log.Flags())
	router := mux.NewRouter()
	h := handlers.Init(database, logger)
	router.HandleFunc("/ping", h.HandlePing)

	logger.Fatal(http.ListenAndServe(address, router))

	return nil
}
