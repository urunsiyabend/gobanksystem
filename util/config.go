package util

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER" env:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE" env:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS" env:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY" env:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_LIFETIME" env:"ACCESS_TOKEN_LIFETIME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
