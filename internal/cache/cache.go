package cache

import (
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"github.com/google/uuid"
)

type Cache interface {
	Start() (Cache, chan database.UpdateMessage, chan database.DeleteMessage)

	Chat(id uuid.UUID) (*database.Chat, bool)
	ChatsByUser(user *database.User) []*database.Chat
	CreateChat(chat *database.Chat) error
	UpdateChat(chat *database.Chat) error
	DeleteChat(id uuid.UUID) error

	Message(id uuid.UUID) (*database.Message, bool)
	CreateMessage(msg *database.Message) error
	UpdateMessage(msg *database.Message) error
	DeleteMessage(id uuid.UUID, deleteForAll bool) error

	User(id uuid.UUID) (*database.User, bool)
	UserByName(username string) (*database.User, bool)
	UsersByNames(usernames []string) []*database.User
	CreateUser(user *database.User) error
	UpdateUser(user *database.User) error
	DeleteUser(id uuid.UUID) error
}

type cache struct {
	log *logger.Logger

	user    map[uuid.UUID]*database.User
	message map[uuid.UUID]*database.Message
	chat    map[uuid.UUID]*database.Chat

	updateChan chan database.UpdateMessage
	deleteChan chan database.DeleteMessage

	getSnapshot func() ([]*database.User, []*database.Message, []*database.Chat)
}

func NewCache(
	logger *logger.Logger,
	snapshotFunc func() ([]*database.User, []*database.Message, []*database.Chat),
) Cache {
	return &cache{
		getSnapshot: snapshotFunc,
		log:         logger,
	}
}
func (c *cache) Start() (Cache, chan database.UpdateMessage, chan database.DeleteMessage) {
	c.convertSnapshot(c.getSnapshot())
	c.updateChan = make(chan database.UpdateMessage)
	c.deleteChan = make(chan database.DeleteMessage)
	return c, c.updateChan, c.deleteChan
}

func (c *cache) convertSnapshot(usr []*database.User, msg []*database.Message, chat []*database.Chat) {
	c.user = convertUsers(c.log, usr)
	c.chat = convertChats(c.log, chat)
	c.message = convertMessages(c.log, msg)
}

func convertUsers(log *logger.Logger, usr []*database.User) map[uuid.UUID]*database.User {
	res := make(map[uuid.UUID]*database.User)
	for _, u := range usr {
		id := u.GetId()
		if id == uuid.Nil {
			id = uuid.New()
			log.Warning("User with name:%s has empty id, new id:%s", u.Name, id)
		}
		res[id] = u
	}
	return res
}

func convertChats(log *logger.Logger, usr []*database.Chat) map[uuid.UUID]*database.Chat {
	res := make(map[uuid.UUID]*database.Chat)
	for _, u := range usr {
		id := u.GetId()
		if id == uuid.Nil {
			id = uuid.New()
			log.Warning("Chat with name:%s has empty id, new id:%s", u.Name, id)
		}
		res[id] = u
	}
	return res
}

func convertMessages(log *logger.Logger, usr []*database.Message) map[uuid.UUID]*database.Message {
	res := make(map[uuid.UUID]*database.Message)
	for _, u := range usr {
		id := u.GetId()
		if id == uuid.Nil {
			id = uuid.New()
			log.Warning("Message has empty id, new id:%s", id)
		}
		res[id] = u
	}
	return res
}
