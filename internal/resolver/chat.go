package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"time"
)

func (r Resolver) ChatWS(ws *websocket.Conn, chatId string) {

	for {
		ws.ReadMessage()
		err := ws.WriteJSON("dfg")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (r Resolver) CreateChat(chat models.AddChat) (*database.Chat, error) {
	res := r.Db.Create(&database.Chat{
		Name:      chat.Name,
		CreatedAt: time.Now(),
		Users:     makeUsersForChat(r.Db, chat.Users),
		Messages:  nil,
	})
	if res.Error != nil {
		return nil, res.Error
	}

	return res.Statement.Model.(*database.Chat), nil
}

func (r Resolver) GetUserChats(userId string) []*database.Chat {
	var chats []*database.Chat
	var res []*database.Chat
	r.Db.
		Preload("Users").
		Find(&chats)

	// TODO сделать фильтауию через запрос
	for _, chat := range chats {
		for _, user := range chat.Users {
			if user.Username == userId {
				res = append(res, chat)
			}
		}
	}
	fmt.Println("find: ", len(res))
	return chats
}

func makeUsersForChat(db *gorm.DB, usernames []string) []*database.User {
	res := make([]*database.User, 0)
	db.Where("username IN ?", usernames).Find(&res)
	return res
}
