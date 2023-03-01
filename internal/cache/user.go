package cache

import (
	"Messenger/internal/database"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (c *cache) CreateUser(user *database.User) error {
	if _, ok := c.user[user.Id]; ok {
		return fmt.Errorf("usr with id:%d already exist", user.Id)
	}
	user.CreatedAt = time.Now()
	c.user[user.Id] = user

	c.updateChan <- UpdateMessage{User: []*database.User{user}}
	return nil
}

func (c *cache) UpdateUser(user *database.User) error {
	if _, ok := c.user[user.Id]; !ok {
		return fmt.Errorf("usr with id:%d does not exist", user.Id)
	}
	c.user[user.Id] = user

	c.updateChan <- UpdateMessage{User: []*database.User{user}}
	return nil
}

func (c *cache) DeleteUser(id uuid.UUID) error {
	usr, ok := c.user[id]
	if !ok {
		return fmt.Errorf("usr with id:%d does not exist", id)
	}

	delete(c.user, id)

	c.deleteChan <- DeleteMessage{User: []uuid.UUID{id}}
	c.updateChan <- UpdateMessage{User: []*database.User{usr}}
	return nil
}

func (c *cache) User(id uuid.UUID) (*database.User, bool) {
	u, ok := c.user[id]
	return u, ok
}
