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

	c.updateChan <- database.UpdateMessage{User: []*database.User{user}}
	return nil
}

func (c *cache) UpdateUser(user *database.User) error {
	if _, ok := c.user[user.Id]; !ok {
		return fmt.Errorf("usr with id:%d does not exist", user.Id)
	}
	c.user[user.Id] = user

	c.updateChan <- database.UpdateMessage{User: []*database.User{user}}
	return nil
}

func (c *cache) DeleteUser(id uuid.UUID) error {
	_, ok := c.user[id]
	if !ok {
		return fmt.Errorf("usr with id:%d does not exist", id)
	}

	delete(c.user, id)

	c.deleteChan <- database.DeleteMessage{User: []uuid.UUID{id}}
	return nil
}

func (c *cache) User(id uuid.UUID) (*database.User, bool) {
	u, ok := c.user[id]
	if !ok {
		return nil, false
	}
	return u, ok
}

func (c *cache) UserByName(username string) (*database.User, bool) {
	for _, user := range c.user {
		if user.Name == username {
			return user, true
		}
	}
	return nil, false
}

func (c *cache) UsersByNames(usernames []string) []*database.User {
	res := make([]*database.User, 0)
	for _, user := range c.user {
		if contains(usernames, user.Name) {
			res = append(res, user)
		}
	}
	return res
}
