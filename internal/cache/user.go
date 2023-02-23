package cache

import (
	"Messenger/database"
	"github.com/google/uuid"
)

func (c *cache) CreateUser(user database.User) bool {
	//TODO implement me
	panic("implement me")
}

func (c *cache) UpdateUser(user database.User) bool {
	//TODO implement me
	panic("implement me")
}

func (c *cache) DeleteUser(id uuid.UUID) bool {
	//TODO implement me
	panic("implement me")
}

func (c *cache) User(id uuid.UUID) database.User {
	//TODO implement me
	panic("implement me")
}
