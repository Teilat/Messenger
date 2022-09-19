package webapi

import (
	_ "Messenger/webapi/docs"
	"Messenger/webapi/globals"
	"Messenger/webapi/handlers"
	"fmt"
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

	cookieStore := cookie.NewStore(globals.Secret)
	cookieStore.Options(sessions.Options{HttpOnly: false, Secure: false, MaxAge: 86400, Path: "/"})
	router.Use(sessions.Sessions("token", cookieStore))
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Authorization", "Origin", "Accept", "X-Requested-With", "Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Access-Control-Allow-Credentials"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))
	logger := log.New(os.Stderr, "Handler: ", log.LstdFlags)
	h := handlers.Init(database, logger)

	router.GET("/", h.HandlePing())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", h.LoginPostHandler())
	router.POST("/register", h.RegisterHandler())
	router.GET("/logout", h.LogoutGetHandler())

	router.GET("/chats", h.GetAllChatsHandler())
	router.POST("/chat", h.CreateChatHandler())
	router.GET("/chat/:id", h.GetChatHandler())
	err := router.Run(address)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
