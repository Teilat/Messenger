package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func CreateUser(db *gorm.DB, user models.AddUser) error {
	res := db.Create(&database.User{
		Username: user.Username,
		Name:     user.Nickname,
		Phone:    user.Phone,
		PwHash:   user.Password,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func GetUserByUsername(db *gorm.DB, username string) *database.User {
	user := database.User{}
	db.Where("username = ?", username).Preload("Chats").First(&user)
	fmt.Println("found:", user)
	updateLastOnline(db, username)
	return &user
}

func updateLastOnline(db *gorm.DB, username string) {
	updateTime := time.Now()
	db.First(&database.User{Username: username}).Update("last_online", updateTime)
}
