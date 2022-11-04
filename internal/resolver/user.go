package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"fmt"
	"mime/multipart"
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
	r.Db.Find(&user.Chats).Preload("Messages")
	r.updateLastOnline(username)
	return &user
}

func (r Resolver) UpdateUserImage(username string, image *multipart.FileHeader) error {
	user := r.GetUserByUsername(username)
	asd, _ := image.Open()
	asd.Read(user.Image)

	res := r.Db.Save(user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r Resolver) updateLastOnline(username string) {
	updateTime := time.Now()
	r.Db.First(&database.User{Username: username}).Update("last_online", updateTime)
}
