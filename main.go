package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("hyperman")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fata error config file: %s \n", err))
	}
}

func main() {
	fmt.Println("I am Hyperman")
}
