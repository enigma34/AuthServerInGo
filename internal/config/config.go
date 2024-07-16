package config

import (
	"github.com/spf13/viper"
)

type JWTConfig struct {
	SigningKey  string `json:"signingKey"`
	Issuer      string `json:"issuer"`
	ExpiryHours int    `json:"expiryHours"`
}

type Config struct {
	JWT JWTConfig `json:"jwt"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
