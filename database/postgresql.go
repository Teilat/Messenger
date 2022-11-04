package database

import (
	"Messenger/internal/logger"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbProvider struct {
	connString string
	log        *logger.MyLog
}

func NewDbProvider(log *logger.MyLog) *DbProvider {
	return &DbProvider{
		connString: fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			viper.Get("postgresql.user"),
			viper.Get("postgresql.pass"),
			viper.Get("postgresql.host"),
			viper.Get("postgresql.port"),
			viper.Get("postgresql.db")),
		log: log}
}

func (p DbProvider) Run() *gorm.DB {
	db, err := gorm.Open(postgres.Open(p.connString), &gorm.Config{})
	if err != nil {
		p.log.Error("error while connecting to database:" + err.Error())
		return nil
	}
	// creating tables from code models
	err = db.AutoMigrate(&User{}, &Chat{}, &Message{})
	if err != nil {
		p.log.Error("error migrating database:" + err.Error())
		return nil
	}
	return db
}
