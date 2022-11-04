package cache

import (
	"Messenger/internal/logger"
	"gorm.io/gorm"
)

type C interface {
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

type Cache struct {
	db  *gorm.DB
	log *logger.MyLog
}

func NewCache(logger *logger.MyLog, db *gorm.DB) *Cache {
	return &Cache{
		db:  db,
		log: logger,
	}
}
func (c Cache) Run() {

}

func (c Cache) CreateChat() {
	//TODO implement me
	panic("implement me")
}

func (c Cache) UpdateChat() {
	//TODO implement me
	panic("implement me")
}

func (c Cache) DeleteChat() {
	//TODO implement me
	panic("implement me")
}

func (c Cache) CreateMessage() {
	//TODO implement me
	panic("implement me")
}

func (c Cache) UpdateMessage() {
	//TODO implement me
	panic("implement me")
}

func (c Cache) DeleteMessage() {
	//TODO implement me
	panic("implement me")
}

func (c Cache) CreateUser() {
	//TODO implement me
	panic("implement me")
}

func (c Cache) UpdateUser() {
	//TODO implement me
	panic("implement me")
}

func (c Cache) DeleteUser() {
	//TODO implement me
	panic("implement me")
}
