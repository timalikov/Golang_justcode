package main

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HttpServer HttpServerConfig
	Database   DatabaseConfig
}

type HttpServerConfig struct {
	Port            int           `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}

type DatabaseConfig struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
}

// Инициализация и чтение конфигурации
func initConfig() (*Config, error) {
	var config Config

	viper.SetConfigType("yaml")
	viper.SetConfigName("config") // Имя файла конфигурации без расширения
	viper.AddConfigPath(".")      // Путь к каталогу с файлом конфигурации

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	_, err := initConfig()
	if err != nil {
		panic(err)
	}

}
