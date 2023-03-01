package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	Postgres Postgresql `yaml:"postgresql"`
	Api      Api        `yaml:"api"`
}

type Api struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type Postgresql struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Database string `yaml:"db"`
	Password string `yaml:"pass"`
}

func GetConf() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}
