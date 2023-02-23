package cache

import (
	"Messenger/database"
	"Messenger/internal/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cache interface {
	Start() (chan UpdateMessage, chan DeleteMessage)
	run()

	Chat(id uuid.UUID) database.Chat
	CreateChat(chat database.Chat) bool
	UpdateChat(chat database.Chat) bool
	DeleteChat(id uuid.UUID) bool

	Message(id uuid.UUID) database.Message
	CreateMessage(msg database.Message) bool
	UpdateMessage(msg database.Message) bool
	DeleteMessage(id uuid.UUID) bool

	User(id uuid.UUID) database.User
	CreateUser(user database.User) bool
	UpdateUser(user database.User) bool
	DeleteUser(id uuid.UUID) bool
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

	user    []*database.User
	message []*database.Message
	chat    []*database.Chat

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
