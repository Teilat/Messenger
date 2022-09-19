package handlers

import (
	"Messenger/database"
	"Messenger/internal/resolver"
	"Messenger/webapi/converters"
	"Messenger/webapi/globals"
	"Messenger/webapi/helpers"
	"Messenger/webapi/models"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Handlers struct {
	log      *log.Logger
	upgrader *websocket.Upgrader
	Resolver *resolver.Resolver
}

func Init(logger *log.Logger, res *resolver.Resolver) *Handlers {
	return &Handlers{
		log: logger,
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		Resolver: res,
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
		session := sessions.Default(c)
		sessionUser := session.Get(globals.Userkey)
		if sessionUser != nil {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Please logout first"})
			return
		}
		var params models.Login

		err := c.BindJSON(&params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to parse params"})
			return
		}

		if helpers.EmptyUserPass(params) {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(h.Resolver.Db, params) {
			c.JSON(http.StatusUnauthorized, gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, params.Username)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to save session"})
			return
		}

		c.JSON(http.StatusOK, converters.UserToApiUser(h.Resolver.GetUserByUsername(params.Username), []*database.Chat{}))
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
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("logging out user:", user)
		if user == nil {
			log.Println("Invalid session token")
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Invalid session token"})
			return
		}
		session.Delete(user)
		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to save session"})
			return
		}
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
		session := sessions.Default(c)
		sessionUser := session.Get(globals.Userkey)
		if sessionUser != nil {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Please logout first"})
			return
		}

		var user models.AddUser
		err := c.BindJSON(&user)
		if err != nil {
			log.Println(err)
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
// @Produce     json
// @Success     101 {object} models.WSChat "ws struct"
// @Error       500 {string} string
// @Router      /chat/:id [post]
func (h Handlers) WSChatHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//upgrade get request to websocket protocol
		var id = c.Param("id")
		ws, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {
				h.log.Fatal(err)
			}
		}(ws)

		h.Resolver.ChatWS(ws, id)
	}
}

// GetAllChatsHandler  godoc
// @Summary     get all chats
// @Tags        Chat
// @Accept      json
// @Produce     json
// @Success     200 {array} models.Chat "list of chats for current user"
// @Error       500 {string} string
// @Router      /chats [get]
func (h Handlers) GetAllChatsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUser := session.Get(globals.Userkey)
		if sessionUser == nil {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Please login first"})
			return
		}
		c.JSON(http.StatusOK, converters.ChatsToApiChats(h.Resolver.GetUserChats(sessionUser.(string))))
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
		session := sessions.Default(c)
		sessionUser := session.Get(globals.Userkey)
		if sessionUser == nil {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Please login first"})
			return
		}

		var params models.AddChat
		err := c.BindJSON(&params)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to parse params"})
		}
		params.Users = append(params.Users, sessionUser.(string))
		chat, err := h.Resolver.CreateChat(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to create chat"})
		}

		c.JSON(http.StatusOK, converters.ChatToApiChat(chat))
	}
}
