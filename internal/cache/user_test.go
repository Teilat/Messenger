package cache

import (
	"Messenger/internal/database"
	"Messenger/internal/logger"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_User(t *testing.T) {
	c, upd, del := NewCache(logger.NewLogger("[Test Cache]"),
		func() ([]*database.User, []*database.Message, []*database.Chat) {
			return []*database.User{
					{Id: uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"), Name: "new user 1"},
				},
				[]*database.Message{},
				[]*database.Chat{{
					Id:   uuid.MustParse("c91d7e4e-a33f-4dd3-91ef-bb8fa9904de5"),
					Name: "chat 1",
					Users: []*database.User{
						{Id: uuid.MustParse("3c54480e-d09a-11ed-afa1-0242ac120002"), Name: "new user 1", Username: "user1"},
					},
				}}
		}).Start()
	database.NewDbProvider(logger.NewLogger("[Test Database]")).StartUpdateListener(context.TODO(), upd, del)

	usr := &database.User{
		Id:         uuid.New(),
		Username:   "new user 2",
		Name:       "user 2",
		Bio:        "bio",
		Phone:      "+78005553535",
		PwHash:     "password",
		CreatedAt:  time.Now(),
		LastOnline: time.Now(),
		Chats:      []*database.Chat{},
		Messages:   []database.Message{},
	}

	// create message
	err := c.CreateUser(usr)
	assert.NoErrorf(t, err, "CreateUser error:%v", err)

	createdUsr, ok := c.User(usr.Id)
	if !ok {
		t.Error("Get user error")
	}
	assert.Equalf(t, usr, createdUsr, "User not equal \nExpected:%+v, \nGot:%+v", usr, createdUsr)

	createdUsrByName, ok := c.UserByName(usr.Username)
	if !ok {
		t.Error("Get user error")
	}
	assert.Equalf(t, usr, createdUsrByName, "User not equal \nExpected:%+v, \nGot:%+v", usr, createdUsrByName)

	// update message
	usr.Bio = "updated text"
	err = c.UpdateUser(usr)
	assert.NoErrorf(t, err, "UpdateUser error:%v", err)

	updatedUsr, ok := c.User(usr.Id)
	if !ok {
		t.Error("Get user error")
	}
	assert.Equalf(t, usr, updatedUsr, "User not equal \nExpected:%+v, \nGot:%+v", usr, updatedUsr)

	// delete message
	err = c.DeleteUser(usr.Id)
	assert.NoErrorf(t, err, "CreateUser error:%v", err)

	deletedUsr, ok := c.User(usr.Id)
	if ok {
		t.Error("Get deleted user error")
	}
	assert.Equalf(t, (*database.User)(nil), deletedUsr, "Chats not equal \nExpected:%+v, \nGot:%+v", nil, deletedUsr)
}
