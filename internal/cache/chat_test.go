package cache

import (
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

// full test for CRUD
func Test_Chat(t *testing.T) {
	c, upd, del := NewCache(logger.NewLogger("[Test Cache]"),
		func() ([]*database.User, []*database.Message, []*database.Chat) {
			return []*database.User{}, []*database.Message{}, []*database.Chat{}
		}).Start()
	database.NewDbProvider(logger.NewLogger("[Test Database]")).StartUpdateListener(context.TODO(), upd, del)

	chat := &database.Chat{
		Id:       uuid.New(),
		Name:     "Test chat",
		Users:    []*database.User{},
		Messages: []database.Message{},
	}
	// create chat
	err := c.CreateChat(chat)
	assert.NoErrorf(t, err, "CreateChat error:%v", err)

	createdChat, ok := c.Chat(chat.Id)
	if !ok {
		t.Error("Get chat error")
	}
	assert.Equalf(t, chat, createdChat, "Chats not equal \nExpected:%+v, \nGot:%+v", chat, createdChat)

	// update chat
	chat.Name = "Updated test chat"
	err = c.UpdateChat(chat)
	assert.NoErrorf(t, err, "UpdateChat error:%v", err)

	updatedChat, ok := c.Chat(chat.Id)
	if !ok {
		t.Error("Get chat error")
	}
	assert.Equalf(t, chat, updatedChat, "Chats not equal \nExpected:%+v, \nGot:%+v", chat, updatedChat)

	// delete chat
	err = c.DeleteChat(chat.Id)
	assert.NoErrorf(t, err, "CreateChat error:%v", err)

	deletedChat, ok := c.Chat(chat.Id)
	if ok {
		t.Error("Get deleted chat error")
	}
	assert.Equalf(t, (*database.Chat)(nil), deletedChat, "Chats not equal \nExpected:%+v, \nGot:%+v", nil, deletedChat)
}

func Test_cache_ChatsByUser(t *testing.T) {
	tests := []struct {
		name   string
		userId uuid.UUID
		chatId uuid.UUID
	}{
		{
			name:   "",
			userId: uuid.MustParse("3c544cbe-d09a-11ed-afa1-0242ac120002"),
			chatId: uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"),
		},
	}
	c := &cache{
		log: logger.NewLogger("[Test Cache]"),
		user: map[uuid.UUID]*database.User{
			uuid.MustParse("3c544cbe-d09a-11ed-afa1-0242ac120002"): {
				Id:   uuid.MustParse("3c544cbe-d09a-11ed-afa1-0242ac120002"),
				Name: "user 1",
			},
		},
		chat: map[uuid.UUID]*database.Chat{
			uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"): {
				Id:   uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"),
				Name: "chat 1",
				Users: []*database.User{
					{
						Id:   uuid.MustParse("3c544cbe-d09a-11ed-afa1-0242ac120002"),
						Name: "user 1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := c.User(tt.userId)
			want, _ := c.Chat(tt.chatId)
			got := c.ChatsByUser(user)
			assert.Equalf(t, []*database.Chat{want}, got, "ChatsByUser(%v)", user)
		})
	}
}
