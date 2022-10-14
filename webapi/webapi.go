package webapi

import (
	"Messenger/internal/logger"
	"Messenger/internal/resolver"
	_ "Messenger/webapi/docs"
	"Messenger/webapi/handlers"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	"os"
)

// @Title     Application Api
// @Version   1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func Run(database *gorm.DB) error {
	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))

	resolverLog := logger.NewLogger(log.New(os.Stderr, "[Resolver] ", log.LstdFlags))
	res := resolver.Init(database, resolverLog)

	hub := resolver.NewHub()
	go hub.Run()

	handlerLog := logger.NewLogger(log.New(os.Stderr, "[Handler] ", log.LstdFlags))
	h := handlers.Init(handlerLog, res, hub)

	authMiddleware, err := jwt.New(newJwtMiddleware(h.Resolver, true))
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("Auth middleware init error:" + errInit.Error())
	}

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(newCors()))

	authGroup := router.Group("")
	authGroup.Use(authMiddleware.MiddlewareFunc())

	router.GET("/", h.HandlePing())
	router.GET("/debug", h.HandleDebug())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/register", h.RegisterHandler())
	authGroup.GET("/logout", authMiddleware.LogoutHandler)
	authGroup.GET("/user/:username", h.GetUser())

	authGroup.GET("/chats", h.GetAllChatsHandler())
	authGroup.POST("/chat", h.CreateChatHandler())
	authGroup.GET("/chat/:id", h.WSChatHandler())

	err = router.Run(address)
	//err := router.RunTLS(address, "./server-cert.pem", "./server-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
