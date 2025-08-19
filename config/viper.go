package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ModelInfo struct {
		ModelName   string  `mapstructure:"model_name"`
		BaseUrl     string  `mapstructure:"base_url"`
		ApiKey      string  `mapstructure:"api_key"`
		Temperature float32 `mapstructure:"temperature"`
	} `mapstructure:"model_info"`
	Memobase struct {
		ProjectURL string `mapstructure:"project_url"`
		APIKey     string `mapstructure:"api_key"`
	} `mapstructure:"memobase"`
}

var Conf Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("unable to decode config, %v", err)
	}
}
