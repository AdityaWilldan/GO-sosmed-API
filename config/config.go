package config

import "github.com/spf13/viper"

type Config struct {
	PORT        string
	DB_USERNAME string
	DB_PASSWORD string
	DB_URL      string
	DB_DATABASE string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	} else {
		viper.Unmarshal(&ENV)
	}
}
