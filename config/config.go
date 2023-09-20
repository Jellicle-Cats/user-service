package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	MongoURL       string        `mapstructure:"MONGO_URL"`
	JWTTokenSecret string        `mapstructure:"JWT_SECRET"`
	ClientID       string        `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	ClientSecret   string        `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	RedirectURL    string        `mapstructure:"REDIRECT_URL"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

func ReadConfig() Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("[Config] Cannot read config file, %s", err))
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("[Config] Unable to decode into struct, %s", err))
	}

	return config
}
