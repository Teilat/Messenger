package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (r Resolver) CreateUser(user models.AddUser) error {
	res := r.Db.Create(&database.User{
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

func (r Resolver) GetUserByUsername(username string) *database.User {
	user := database.User{}
	r.Db.Where("username = ?", username).Preload("Chats").First(&user)
	fmt.Println("found:", user.Username)
	updateLastOnline(r.Db, username)
	return &user
}

func updateLastOnline(db *gorm.DB, username string) {
	updateTime := time.Now()
	db.First(&database.User{Username: username}).Update("last_online", updateTime)
}
