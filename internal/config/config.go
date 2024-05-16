package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Port    int           `yaml:"port" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"10s"`
	AdminId []int         `yaml:"admin_id" env-require:"true"`
}

func MustLoad() *Config {
	var path string
	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parse()
	if path == "" {
		panic("empty path")
	}
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
