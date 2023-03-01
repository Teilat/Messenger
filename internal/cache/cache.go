package cache

import (
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cache interface {
	Start() (chan UpdateMessage, chan DeleteMessage)

	Chat(id uuid.UUID) (*database.Chat, bool)
	CreateChat(chat *database.Chat) error
	UpdateChat(chat *database.Chat) error
	DeleteChat(id uuid.UUID) error

	Message(id uuid.UUID) (*database.Message, bool)
	CreateMessage(msg *database.Message) error
	UpdateMessage(msg *database.Message) error
	DeleteMessage(id uuid.UUID, deleteForAll bool) error

	User(id uuid.UUID) (*database.User, bool)
	CreateUser(user *database.User) error
	UpdateUser(user *database.User) error
	DeleteUser(id uuid.UUID) error
}

type UpdateMessage struct {
	User    []*database.User
	Message []*database.Message
	Chat    []*database.Chat
}

type DeleteMessage struct {
	User    []uuid.UUID
	Message []uuid.UUID
	Chat    []uuid.UUID
}

type cache struct {
	db  *gorm.DB
	log *logger.Logger

	user    map[uuid.UUID]*database.User
	message map[uuid.UUID]*database.Message
	chat    map[uuid.UUID]*database.Chat

	updateChan chan UpdateMessage
	deleteChan chan DeleteMessage
}

func NewCache(logger *logger.Logger, db *gorm.DB) Cache {
	return &cache{
		db:  db,
		log: logger,
	}
}
func (c *cache) Start() (chan UpdateMessage, chan DeleteMessage) {
	go c.run()
	return c.updateChan, c.deleteChan
}

func (c *cache) run() {
	c.updateChan = make(chan UpdateMessage)
	c.deleteChan = make(chan DeleteMessage)
	for {
		select {}
	}
}
