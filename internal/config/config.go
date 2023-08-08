package config

import (
	"fmt"
	"sync"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Http struct {
		Ip          string  `yaml:"ip" env-required:"true"`
		Port        int     `yaml:"port" env-required:"true"`
		ReadTimeout int `yaml:"read-timeout" env-required:"true"`
		WriteTimeout int `yaml:"write-timeout" env-required:"true"`
    CORS         struct {
			AllowedMethods     []string `yaml:"allowed-methods" env-required:"true"`
			AllowedOrigins     []string `yaml:"allowed-origins" env-required:"true"`
			AllowCredentials   bool     `yaml:"allow-credentials"`
			AllowedHeaders     []string `yaml:"allowed-headers" env-required:"true"`
			ExposedHeaders     []string `yaml:"exposed-headers" env-required:"true"`
			Debug              bool     `yaml:"debug"`
		} `yaml:"cors"`
	} `yaml:"http" env-required:"true"`
}

var(
  config = &Config{}
  once = sync.Once{}
)

func GetConfig() *Config {
  once.Do(func (){
    path := "configs/config.yml"
    config = &Config{}

    if err := cleanenv.ReadConfig(path, config); err != nil {
      panic(fmt.Sprintf("Failed to read config in path: %s error: %s", path, err))
    }
  })
  
  return config
}
