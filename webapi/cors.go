package webapi

import (
	"github.com/gin-contrib/cors"
	"time"
)

func newCors() cors.Config {
	return cors.Config{
		ExposeHeaders:    []string{"Access-Token", "Expire-Token"},
		AllowOrigins:     []string{"http://192.168.1.44:8080", "http://192.168.1.134:3000", "http://localhost:3000", "http://192.168.1.1:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE"},
		AllowHeaders:     []string{"jwt", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Access-Control-Allow-Credentials", "Authorization", "Origin", "Accept", "X-Requested-With", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
		AllowWebSockets:  true,
	}
}
