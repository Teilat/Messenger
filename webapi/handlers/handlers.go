package handlers

import (
	"Messenger/database"
	"Messenger/internal/resolver"
	"Messenger/webapi/converters"
	"Messenger/webapi/globals"
	"Messenger/webapi/helpers"
	"Messenger/webapi/models"
	"github.com/gin-contrib/sessions"
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
func (h handlers) LoginPostHandler() gin.HandlerFunc {
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
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if helpers.EmptyUserPass(params) {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(h.db, params) {
			c.JSON(http.StatusUnauthorized, gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, params.Username)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to save session"})
			return
		}

		c.JSON(http.StatusOK, converters.UserToApiUser(resolver.GetUserByUsername(h.db, params.Username)))
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
func (h handlers) LogoutGetHandler() gin.HandlerFunc {
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
func (h handlers) RegisterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUser := session.Get(globals.Userkey)
		if sessionUser != nil {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Please logout first"})
			return
		}

		var params models.AddUser
		err := c.BindJSON(&params)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to register"})
		}

		res := h.db.Create(&database.User{
			Username: params.Username,
			Name:     params.Nickname,
			Phone:    params.Phone,
			PwHash:   params.Password,
		})
		if res.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to register"})
		}
	}
}
