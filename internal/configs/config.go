package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func Initiator() {
	viper.SetConfigFile(".env") 
	viper.SetConfigType("env")  
	viper.AutomaticEnv()       

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading .env file:", err)
		return
	}

	fmt.Println("Successfully read .env file")
}