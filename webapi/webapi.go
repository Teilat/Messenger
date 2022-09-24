package webapi

import (
	"Messenger/internal/resolver"
	_ "Messenger/webapi/docs"
	"Messenger/webapi/globals"
	"Messenger/webapi/handlers"
	"fmt"
	session "github.com/ScottHuangZL/gin-jwt-session"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// @Title     Application Api
// @Version   1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name token

func Run(database *gorm.DB) error {
	// swag init --parseDependency --parseInternal -g webapi.go
	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	session.NewStore()
	cookieStore := cookie.NewStore(globals.Secret)
	cookieStore.Options(sessions.Options{HttpOnly: false, Secure: false, MaxAge: 86400, Path: "/", Domain: "localhost"})
	router.Use(session.ClearMiddleware()) //important to avoid mem leak
	router.Use(sessions.Sessions("token", cookieStore))
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://192.168.1.44:8080", "http://localhost:3000", "http://192.168.1.1:8080"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Access-Control-Allow-Credentials", "Authorization", "Origin", "Accept", "X-Requested-With", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
		AllowWebSockets:  true,
	}))
	logger := log.New(os.Stderr, "Handler: ", log.LstdFlags)
	res := resolver.Init(database, logger)
	h := handlers.Init(logger, res)

	router.GET("/", h.HandlePing())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", h.LoginPostHandler())
	router.POST("/register", h.RegisterHandler())
	router.GET("/logout", h.LogoutGetHandler())

	router.GET("/chats", h.GetAllChatsHandler())
	router.POST("/chat", h.CreateChatHandler())
	router.PATCH("/chat", h.CreateChatHandler())
	router.POST("/chat/:id", h.WSChatHandler())

	router.POST("/message", h.CreateMessageHandler())
	router.PATCH("/message", h.EditMessageHandler())
	router.DELETE("/message", h.DeleteMessageHandler())
	router.GET("/message/:offset", h.GetMessagesBatchHandler())

	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
