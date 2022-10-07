package handlers

import (
	"Messenger/internal/logger"
	"Messenger/internal/resolver"
	"Messenger/webapi/converters"
	"Messenger/webapi/models"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"runtime"
)

type Handlers struct {
	log       *logger.Log
	upgrader  *websocket.Upgrader
	Resolver  *resolver.Resolver
	hub       *resolver.Hub
	LoginUser string
}

func Init(log *logger.Log, res *resolver.Resolver, hub *resolver.Hub) *Handlers {
	return &Handlers{
		log: log,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Resolver: res,
		hub:      hub,
	}
}

// HandlePing   godoc
// @Summary		Health check
// @Tags        General
// @Accept      json
// @Produce     json
// @Success     200 {string} string "healthy"
// @Router      / [get]
func (h Handlers) HandlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK v1")
	}
}

// LoginPostHandler  godoc
// @Summary     Login user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       credentials body models.Login true "credentials"
// @Success     200 {object} models.User "logged in user"
// @Error       500 {string} string
// @Error       404 {string} string
// @Router      /login [post]
func (h Handlers) LoginPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// LogoutGetHandler  godoc
// @Summary     Logout user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Success     200
// @Error       500 {string} string
// @Router      /logout [get]
func (h Handlers) LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// RegisterHandler  godoc
// @Summary     register user
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       user body models.AddUser true "user"
// @Success     200
// @Error       500 {string} string
// @Router      /register [post]
func (h Handlers) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.AddUser
		err := c.BindJSON(&user)
		if err != nil {
			h.log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to parse params"})
		}

		err = h.Resolver.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to register"})
		}
	}
}

// WSChatHandler  godoc
// @Summary     upgrade request to ws
// @Tags        Chat
// @Accept      json
// @Param       ws struct body models.WsMessage false "ws struct"
// @Produce     json
// @Success     101
// @Error       500 {string} string
// @Router      /chat/{id} [get]
func (h Handlers) WSChatHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		//upgrade get request to websocket protocol
		var id = c.Param("id")
		h.log.Printf("ws for %s", id)
		ws, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		h.Resolver.ServeWs(h.hub, ws, h.Resolver, claims[jwt.IdentityKey].(string), id)
	}
}

// GetAllChatsHandler  godoc
// @Summary     Get all chats
// @Tags        Chat
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Chat "list of chats for current user"
// @Error       500 {string} string
// @Router      /chats [get]
func (h Handlers) GetAllChatsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		c.JSON(http.StatusOK, converters.ChatsToApiChats(h.Resolver.GetUserChats(claims[jwt.IdentityKey].(string))))
	}
}

// CreateChatHandler  godoc
// @Summary     create chat
// @Tags        Chat
// @Accept      json
// @Param       chat body models.AddChat true "chat params"
// @Produce     json
// @Success     200 {object} models.Chat "created chat"
// @Error       500 {string} string
// @Router      /chat [post]
func (h Handlers) CreateChatHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		var params models.AddChat
		err := c.BindJSON(&params)
		if err != nil {
			h.log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to parse params"})
		}
		params.Users = append(params.Users)
		chat, err := h.Resolver.CreateChat(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to create chat"})
		}

		c.JSON(http.StatusOK, converters.ChatToApiChat(chat))
	}
}

func (h Handlers) UpdateChatHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// GetUser  godoc
// @Summary     get user by username
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       username path string true "username"
// @Success     200 {object} models.User
// @Error       500 {string} string
// @Router      /user/{username} [get]
func (h Handlers) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, converters.UserToApiUser(h.Resolver.GetUserByUsername(c.Param("username"))))
	}
}

func (h Handlers) HandleDebug() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"goroutines": runtime.NumGoroutine()})
	}
}
