package helpers

import (
	"Messenger/internal/cache"
	"Messenger/webapi/models"
	"strings"
)

func CheckUserPass(c cache.Cache, credentials models.Login) bool {
	user, ok := c.UserByName(credentials.Username)
	if !ok {
		return false
	}
	if EmptyCredentials(credentials) {
		return false
	}
	return user.PwHash == credentials.Password
}

func EmptyCredentials(credentials models.Login) bool {
	return strings.Trim(credentials.Username, " ") == "" || strings.Trim(credentials.Password, " ") == ""
}
