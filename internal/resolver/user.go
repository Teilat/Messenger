package resolver

import (
	"Messenger/internal/database"
	"Messenger/webapi/models"
	"github.com/google/uuid"
	"mime/multipart"
	"time"
)

func (r Resolver) CreateUser(user models.AddUser) error {
	err := r.Cache.CreateUser(&database.User{
		Id:       uuid.New(),
		Username: user.Username,
		Name:     user.Nickname,
		Phone:    user.Phone,
		PwHash:   user.Password,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r Resolver) GetUserByUsername(username string) *database.User {
	user, ok := r.Cache.UserByName(username)
	if !ok {
		return nil
	}
	r.updateLastOnline(username)
	return user
}

func (r Resolver) UpdateUserImage(username string, image *multipart.FileHeader) error {
	user := r.GetUserByUsername(username)
	// img, _ := image.Open()
	// img.Read(user.Image)

	err := r.Cache.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (r Resolver) updateLastOnline(username string) {
	usr, _ := r.Cache.UserByName(username)
	usr.LastOnline = time.Now()
	r.Cache.UpdateUser(usr)
}
