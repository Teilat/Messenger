package helpers

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strings"
)

func CheckUserPass(db *gorm.DB, credentials models.Login) bool {
	user := database.User{Username: credentials.Username}
	res := db.Find(&user)
	if res.Error != nil {
		fmt.Printf("Cant find user error:%s", res.Error.Error())
	}

	log.Println("checkUserPass", credentials.Username, credentials.Password)

	return user.PwHash == credentials.Password
}

func EmptyUserPass(credentials models.Login) bool {
	return strings.Trim(credentials.Username, " ") == "" || strings.Trim(credentials.Password, " ") == ""
}
