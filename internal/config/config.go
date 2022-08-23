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
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}
