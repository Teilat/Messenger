package resolver

import (
	"Messenger/database"
	"Messenger/webapi/models"
	"sort"
	"time"
)

func (r Resolver) CreateMessage(msg models.AddMessage) error {
	chats := []*database.Chat{}
	r.Db.Where("username = ?", msg.User).Preload("Users").Find(&chats)
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].CreatedAt.Before(chats[j].CreatedAt)
	})
	res := r.Db.Create(&database.Message{
		Text:      msg.Text,
		CreatedAt: time.Now(),
		Username:  msg.User,
		ChatId:    chats[msg.ChatId].Id,
	})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
