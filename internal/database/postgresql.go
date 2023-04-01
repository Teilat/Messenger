package database

import (
	"Messenger/internal/logger"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	Database *gorm.DB

	log *logger.Logger

	connString string
	retryCount int
}

func NewDbProvider(log *logger.Logger) *Database {
	return &Database{
		connString: fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			viper.Get("postgresql.user"),
			viper.Get("postgresql.pass"),
			viper.Get("postgresql.host"),
			viper.Get("postgresql.port"),
			viper.Get("postgresql.db")),
		log: log}
}

func (db *Database) Start() (*Database, error) {
	var err error

	db.log.Info("Connecting")
	db.Database, err = gorm.Open(postgres.Open(db.connString), &gorm.Config{})
	for i := 1; err != nil || i > db.retryCount; i++ {
		db.log.Error("Error while connecting error:%#v \nRetry №%d in %d seconds.\n", err.Error(), i, i*2)
		err = nil

		time.Sleep(time.Second * time.Duration(i*2))
		db.Database, err = gorm.Open(postgres.Open(db.connString), &gorm.Config{})
	}

	if err != nil {
		return nil, errors.New("error while connecting to database:" + err.Error())
	}
	db.log.Info("Connected")
	// creating tables from code models
	err = db.Database.AutoMigrate(&User{}, &Chat{}, &Message{})
	if err != nil {
		return nil, errors.New("error migrating database:" + err.Error())
	}

	return db, nil
}

func (db *Database) GetSnapshot() ([]*User, []*Message, []*Chat) {
	var msg []*Message
	var chat []*Chat
	var usr []*User

	db.Database.Find(&msg)
	db.log.Info("Loaded %d messages into snapshot", len(msg))

	db.Database.Preload("Messages").Find(&chat)
	db.log.Info("Loaded %d chats into snapshot", len(chat))

	db.Database.Preload("Messages").Preload("Chats").Find(&usr)
	db.log.Info("Loaded %d users into snapshot", len(usr))

	return usr, msg, chat
}
