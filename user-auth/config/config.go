package config

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	DATABASE      string `mapstructure:"DATABASE"`
	AUTH_SVC_PORT string `mapstructure:"AUTH_SVC_PORT"`
	USER_SVC_URL  string `mapstructure:"USER_SVC_URL"`
	JWT           string `mapstructure:"SECRET_KEY"`
	AUTHTOKEN     string `mapstructure:"TWILIO_AUTH_TOKEN"`
	ACCOUNTSID    string `mapstructure:"TWILIO_ACCOUNT_SID"`
	SERVICESID    string `mapstructure:"TWILIO_SERVICES_ID"`
}

// var envVariables = []string{
// 	"DATABASE", "SECRET_KEY", "TWILIO_AUTH_TOKEN", "TWILIO_ACCOUNT_SID",
// 	"TWILIO_SERVICES_ID", "AUTH_SVC_URL",
// }

var config Config

func LoadConfig() (config *Config, err error) {

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env") // set the file name and path
	err = viper.ReadInConfig()  // read the config file
	if err != nil {             // handle errors while reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// for _, env := range envVariables {
	// 	if err := viper.BindEnv(env); err != nil {
	// 		return config, err
	// 	}
	// }
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
