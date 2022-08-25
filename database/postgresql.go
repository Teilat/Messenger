package database

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresql() (*gorm.DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		viper.Get("postgresql.user"),
		viper.Get("postgresql.pass"),
		viper.Get("postgresql.host"),
		viper.Get("postgresql.port"),
		viper.Get("postgresql.db"))
	fmt.Printf("--> Connecting to:%s\n", connString)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&User{}, &Chat{}, &Message{})
	if err != nil {
		return nil, errors.New("error migrating database:" + err.Error())
	}
	return db, nil
}
