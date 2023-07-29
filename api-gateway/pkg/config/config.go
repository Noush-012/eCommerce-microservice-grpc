package config

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	AUTH_SVC_URL    string `mapstructure:"AUTH_SVC_URL"`
	PRODUCT_SVC_URL string `mapstructure:"PRODUCT_SVC_URL"`
	USER_SVC_URL    string `mapstructure:"USER_SVC_URL"`
	CART_SVC_URL    string `mapstructure:"CART_SVC_URL"`
	ORDER_SVC_URL   string `mapstructure:"ORDER_SVC_URL"`
	PORT            string `mapstructure:"PORT"`
}

var config Config

func LoadConfig() (Config, error) {

	viper.SetConfigType("env")  // set the file type
	viper.SetConfigFile(".env") // set the file name and path
	err := viper.ReadInConfig() // read the config file
	if err != nil {             // handle errors while reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(config); err != nil {
		return config, err
	}
	return config, nil
}

func GetConfig() Config {
	return config
}

// to get the secret code for JWT
// func GetJWTConfig() string {

// 	return config.JWT
// }
