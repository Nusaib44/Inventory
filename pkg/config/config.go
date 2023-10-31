package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (config Config, err error) {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")

	config.Port = viper.GetString("PORT")
	config.DBUrl = viper.GetString("DB_URL")

	if config.Port == "" {
		config.Port = ":2020"
	}

	fmt.Println(config)

	err = viper.Unmarshal(&config)

	fmt.Println(config)

	if err != nil {
		fmt.Println("error occured the error is :", err.Error())
		println()
		return Config{}, err
	}

	return config, err
}
func LoadENV() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
