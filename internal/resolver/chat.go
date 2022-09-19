package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"time"
)

func ChatWS(db *gorm.DB, ws *websocket.Conn, id string) {
	for {
		err := ws.WriteJSON("dfg")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func CreateChat(db *gorm.DB, chat models.AddChat) (*database.Chat, error) {
	res := db.Create(&database.Chat{
		Name:      chat.Name,
		CreatedAt: time.Now(),
		Users:     makeUsersForChat(db, chat.Users),
		Messages:  nil,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*database.Chat), nil
}

func makeUsersForChat(db *gorm.DB, usernames []string) []*database.User {
	res := make([]*database.User, 0)
	db.Where("username IN ?", usernames).Find(&res)
	return res
}

func GetUserChats(db *gorm.DB, userId string) []*database.Chat {
	var chats []*database.Chat
	db.Preload("Users").Where("users IN ?", []string{userId}).Find(&chats)
	fmt.Println("find: ", chats)
	return chats
}
