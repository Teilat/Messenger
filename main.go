package main

import (
	"Messenger/internal/config"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// getting config from config file
	config.GetConf()
	for _, s := range viper.AllKeys() {
		fmt.Printf("%s\n", s)
	}
}
