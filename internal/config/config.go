package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	Postgres Postgresql `yaml:"Postgresql"`
	Api      Api        `yaml:"Api"`
}

type Api struct {
	Address string `yaml:"Address"`
	Port    int    `yaml:"Port"`
}

type Postgresql struct {
	Port     int    `yaml:"Port"`
	Host     string `yaml:"Host"`
	User     string `yaml:"User"`
	Database string `yaml:"Db"`
	Password string `yaml:"Pass"`
}

func GetConf() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	setDefaults()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func setDefaults() {
	viper.SetDefault("api.address", "0.0.0.0")
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("postgresql.port", "5432")
	viper.SetDefault("postgresql.host", "0.0.0.0")
	viper.SetDefault("postgresql.user", "postgres")
	viper.SetDefault("postgresql.pass", "postgrespw")
	viper.SetDefault("postgresql.db", "mainDb")
}
