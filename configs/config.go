package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"HOST"`
		User     string `mapstructure:"USER"`
		Password string `mapstructure:"PASSWORD"`
		Name     string `mapstructure:"NAME"`
		Port     string `mapstructure:"PORT"`
	} `mapstructure:"DB"`

	Server struct {
		Port string `mapstructure:"PORT"`
		Mode string `mapstructure:"MODE"`
	} `mapstructure:"SERVER"`

	JWTSecret string `mapstructure:"JWT_SECRET"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile("config.yaml")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Panic().Err(err).Msg("[GetConfig] Error reading config file")
			return
		}

		cfg = &Config{}
		if err := viper.Unmarshal(cfg); err != nil {
			log.Panic().Err(err).Msg("[GetConfig] Error unmarshalling config")
			return
		}
	})

	return cfg
}
