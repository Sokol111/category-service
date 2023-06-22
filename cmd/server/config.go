package server

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Grpc Grpc `mapstructure:"grpc"`
}

type Grpc struct {
	Server Srv `mapstructure:"server"`
}

type Srv struct {
	Port int `mapstructure:"port"`
}

func LoadConfig(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("Failed to unmarshal config: ", err)
	}
	return c
}
