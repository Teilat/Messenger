package helpers

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func CheckUserPass(db *gorm.DB, credentials models.Login) bool {
	user := database.User{}
	res := db.Where("username = ?", credentials.Username).First(&user)
	if res.Error != nil {
		fmt.Printf("Cant find user error:%s", res.Error.Error())
	}

	return user.PwHash == credentials.Password
}

func EmptyUserPass(credentials models.Login) bool {
	return strings.Trim(credentials.Username, " ") == "" || strings.Trim(credentials.Password, " ") == ""
}
