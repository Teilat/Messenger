package cache

import (
	"Messenger/internal/database"
	"github.com/google/uuid"
)

type mockCache struct {
	user    []*database.User
	message []*database.Message
	chat    []*database.Chat
}

func (m mockCache) Start() (chan UpdateMessage, chan DeleteMessage) {
	panic("implement me")
}

func (m mockCache) Chat(id uuid.UUID) (*database.Chat, bool) {
	panic("implement me")
}

func (m mockCache) CreateChat(chat *database.Chat) error {
	panic("implement me")
}

func (m mockCache) UpdateChat(chat *database.Chat) error {
	panic("implement me")
}

func (m mockCache) DeleteChat(id uuid.UUID) error {
	panic("implement me")
}

func (m mockCache) Message(id uuid.UUID) (*database.Message, bool) {
	panic("implement me")
}

func (m mockCache) CreateMessage(msg *database.Message) error {
	panic("implement me")
}

func (m mockCache) UpdateMessage(msg *database.Message) error {
	panic("implement me")
}

func (m mockCache) DeleteMessage(id uuid.UUID, deleteForAll bool) error {
	panic("implement me")
}

func (m mockCache) User(id uuid.UUID) (*database.User, bool) {
	panic("implement me")
}

func (m mockCache) CreateUser(user *database.User) error {
	panic("implement me")
}

func (m mockCache) UpdateUser(user *database.User) error {
	panic("implement me")
}

func (m mockCache) DeleteUser(id uuid.UUID) error {
	panic("implement me")
}

func NewMockCache() Cache {
	return &mockCache{
		user:    make([]*database.User, 0),
		message: make([]*database.Message, 0),
		chat:    make([]*database.Chat, 0),
	}
}
