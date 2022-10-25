package helpers

import (
	"Messenger/db"
	"Messenger/webapi/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func CheckUserPass(database *gorm.DB, credentials models.Login) bool {
	user := db.User{}
	res := database.Where("username = ?", credentials.Username).First(&user)
	if res.Error != nil {
		fmt.Printf("Cant find user error:%s", res.Error.Error())
	}

	return user.PwHash == credentials.Password
}

func EmptyUserPass(credentials models.Login) bool {
	return strings.Trim(credentials.Username, " ") == "" || strings.Trim(credentials.Password, " ") == ""
}
