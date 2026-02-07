package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	// "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string   `yaml:"env"`
	Database Database `yaml:"database"`
	GRPC     GRPC     `yaml:"grpc"`
}

type GRPC struct {
	Port string `yaml:"port"`
}

type Database struct {
	Port       string `yaml:"port"`
	Host       string `yaml:"host"`
	Name       string `yaml:"name"`
	User       string `yaml:"postgres"`
	Password   string `yaml:"password"`
	ConnString string `yaml:"conn_string"`
}

func Load() (*Config, error) {
	return LoadFrom("config.yaml")
}

func LoadFrom(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	c.Database.ConnString = fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)

	return &c, nil
}

func MustLoad() *Config {
	c, err := Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
	return c
}
