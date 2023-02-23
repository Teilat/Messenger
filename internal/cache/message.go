package cache

import (
	"Messenger/database"
	"github.com/google/uuid"
)

func (c *cache) CreateMessage(msg database.Message) bool {
	//TODO implement me
	panic("implement me")
}

func (c *cache) UpdateMessage(msg database.Message) bool {
	//TODO implement me
	panic("implement me")
}

func (c *cache) DeleteMessage(id uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}

func (c *cache) Message(id uuid.UUID) database.Message {
	//TODO implement me
	panic("implement me")
}
