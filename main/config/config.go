package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBSource          string `mapstructure:"DB_SOURCE"`
	DBName            string `mapstructure:"DB_NAME"`
	LogLevel          string `mapstructure:"LOG_LEVEL"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	CORSAllowOrigins  string `mapstructure:"CORS_ALLOW_ORIGINS"`
	CSRFEnabled       bool   `mapstructure:"CSRF_ENABLED"`
}

func LoadConfig(path, env string) (*Config, error) {
	var config *Config
	viper.AddConfigPath(path)
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}
