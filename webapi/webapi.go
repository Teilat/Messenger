package webapi

import (
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
)

type WebapiProvider struct {
	address  string
	database *gorm.DB
	resolver *resolver.Resolver
	hub      *resolver.Hub
	handler  *handlers.Handlers
}

func NewWebapi(db *gorm.DB, res *resolver.Resolver, hub *resolver.Hub, handlers *handlers.Handlers) *WebapiProvider {
	return &WebapiProvider{
		address:  fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port")),
		database: db,
		resolver: res,
		hub:      hub,
		handler:  handlers,
	}
}

// @Title     Application Api
// @Version   1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func (w WebapiProvider) Run() error {
	go w.hub.Run()

	authMiddleware, err := jwt.New(newJwtMiddleware(w.handler.Resolver, true))
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

	router.GET("/", w.handler.HandlePing())
	router.GET("/debug", w.handler.HandleDebug())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/register", w.handler.RegisterHandler())
	authGroup.GET("/logout", authMiddleware.LogoutHandler)
	authGroup.GET("/user/:username", w.handler.GetUser())

	authGroup.GET("/chats", w.handler.GetAllChatsHandler())
	authGroup.POST("/chat", w.handler.CreateChatHandler())
	authGroup.GET("/chat/:id", w.handler.WSChatHandler())

	err = router.Run(w.address)
	//err := router.RunTLS(address, "./server-cert.pem", "./server-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
