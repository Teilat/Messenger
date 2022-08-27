package handlers

import (
	"github.com/gin-gonic/gin"
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

// HealthCheck  godoc
// @Summary		Health check
// @Tags        General
// @Accept      json
// @Produce     json
// @Success     200 {string} string "healthy"
// @Router      / [get]
func (h handlers) HandlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK v1")
	}
}
