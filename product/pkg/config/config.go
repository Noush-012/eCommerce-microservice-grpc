package config

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	DATABASE         string `mapstructure:"DATABASE"`
	PRODUCT_SVC_PORT string `mapstructure:"USER_SVC_PORT"`
}

func LoadConfig() (config *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env") // set the file name and path
	err = viper.ReadInConfig()  // read the config file
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

// func GetConfig() Config {
// 	var Config Config
// 	return Config
// }
