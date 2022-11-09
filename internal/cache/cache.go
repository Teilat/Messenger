package cache

import (
	"Messenger/database"
	"Messenger/internal/logger"
	"gorm.io/gorm"
)

type Cache interface {
	CreateChat()
	UpdateChat()
	DeleteChat()

	CreateMessage()
	UpdateMessage()
	DeleteMessage()

	CreateUser()
	UpdateUser()
	DeleteUser()
}

type dbStructs struct {
	User    []*database.User
	Message []*database.Message
	Chat    []*database.Chat
}

type cache struct {
	db  *gorm.DB
	log *logger.MyLog

	createChan chan dbStructs
	updateChan chan dbStructs
	deleteChan chan dbStructs
}

func NewCache(logger *logger.MyLog, db *gorm.DB) Cache {
	return &cache{
		db:         db,
		log:        logger,
		createChan: make(chan dbStructs),
		updateChan: make(chan dbStructs),
		deleteChan: make(chan dbStructs),
	}
}
func (c cache) Run() {

}

func (c cache) CreateChat() {
	//TODO implement me
	panic("implement me")
}

func (c cache) UpdateChat() {
	//TODO implement me
	panic("implement me")
}

func (c cache) DeleteChat() {
	//TODO implement me
	panic("implement me")
}

func (c cache) CreateMessage() {
	//TODO implement me
	panic("implement me")
}

func (c cache) UpdateMessage() {
	//TODO implement me
	panic("implement me")
}

func (c cache) DeleteMessage() {
	//TODO implement me
	panic("implement me")
}

func (c cache) CreateUser() {
	//TODO implement me
	panic("implement me")
}

func (c cache) UpdateUser() {
	//TODO implement me
	panic("implement me")
}

func (c cache) DeleteUser() {
	//TODO implement me
	panic("implement me")
}
