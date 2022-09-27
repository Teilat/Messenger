package webapi

import (
	"Messenger/internal/resolver"
	"Messenger/webapi/converters"
	_ "Messenger/webapi/docs"
	"Messenger/webapi/handlers"
	"Messenger/webapi/helpers"
	"Messenger/webapi/models"
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
	"time"
)

// @Title     Application Api
// @Version   1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func Run(database *gorm.DB) error {
	// swag init --parseDependency --parseInternal -g webapi.go
	address := fmt.Sprintf("%s:%d", viper.Get("api.address"), viper.Get("api.port"))

	logger := log.New(os.Stderr, "[Handler] ", log.LstdFlags)
	res := resolver.Init(database, logger)
	hub := resolver.NewHub()
	go hub.Run()
	h := handlers.Init(logger, res, hub)

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		SendCookie:  false,
		CookieName:  "jwt",
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TimeFunc:    time.Now,
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		// --------------------
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if !helpers.CheckUserPass(h.Resolver.Db, loginVals) {
				return "", jwt.ErrFailedAuthentication
			}
			return &models.Login{
				Username: loginVals.Username,
			}, nil
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims { // структура внутри jwt
			if v, ok := data.(*models.Login); ok {
				h.LoginUser = v.Username
				return jwt.MapClaims{
					jwt.IdentityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		LoginResponse: func(c *gin.Context, code int, message string, time time.Time) {
			c.Writer.Header().Add("Access-Token", message)
			c.Writer.Header().Add("Expire-Token", time.Format("2006-01-02 15:04:05"))
			c.JSON(code, converters.UserToApiUser(h.Resolver.GetUserByUsername(h.LoginUser), h.Resolver.GetUserChats(h.LoginUser)))
		},
		//----------------------
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.Login{
				Username: claims[jwt.IdentityKey].(string),
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(code, "")
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.New(cors.Config{
		ExposeHeaders:    []string{"Access-Token", "Expire-Token"},
		AllowOrigins:     []string{"http://192.168.1.44:8080", "http://192.168.1.134:3000", "http://localhost:3000", "http://192.168.1.1:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE"},
		AllowHeaders:     []string{"jwt", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Access-Control-Allow-Credentials", "Authorization", "Origin", "Accept", "X-Requested-With", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
		AllowWebSockets:  true,
	}))

	authGroup := router.Group("")
	authGroup.Use(authMiddleware.MiddlewareFunc())

	router.GET("/", h.HandlePing())
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
