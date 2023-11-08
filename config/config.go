package config

import (
	"os"
	"time"
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
	return Config{
		MongoURL:       os.Getenv("MONGO_URL"),
		JWTTokenSecret: os.Getenv("JWT_SECRET"),
		ClientID:       os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret:   os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		RedirectURL:    os.Getenv("REDIRECT_URL"),
		TokenExpiresIn: 1 * time.Hour,
		TokenMaxAge:    60,
	}
}
