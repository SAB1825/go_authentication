package config

import (
	"log/slog"
	"strings"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/joho/godotenv/autoload"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	HTTP HttpConfig `koanf:"http" validate:"required"`
	DB   DBConfig   `koanf:"db" validate:"required"`
	JWT  JwtConfig  `koanf:"jwt" validate:"required"`
}

type HttpConfig struct {
	Addr         string        `koanf:"addr" validate:"required"`
	ReadTimeout  time.Duration `koanf:"read_timeout" validate:"required"`
	WriteTimeout time.Duration `koanf:"write_timeout" validate:"required"`
	IdleTimeout  time.Duration `koanf:"idle_timeout" validate:"required"`
}

type DBConfig struct {
	URL string `koanf:"url" validate:"required"`
}

type JwtConfig struct {
	SecretKey  string        `koanf:"secret_key" validate:"required"`
	AccessTTL  time.Duration `koanf:"access_ttl" validate:"required"`
	RefreshTTL time.Duration `koanf:"refresh_ttl" validate:"required"`
}

func LoadConfig() (*Config, error) {

	k := koanf.New(".")

	//Loads the env variables and convert them into something like this: GO_AUTH_HTTP_ADDR => http.addr
	err := k.Load(env.Provider("GO_AUTH_", ".", func(s string) string {
		s = strings.TrimPrefix(s, "GO_AUTH_")
		s = strings.ToLower(s)
		return strings.ReplaceAll(s, "__", ".")
	}), nil)
	if err != nil {
		slog.Error("Unable to load the env", "error", err)
		return nil, err
	}

	config := &Config{}

	//Mapping the env variables to the config struct
	err = k.Unmarshal("", config)

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		slog.Error("Failed to validate the struct", "error", err.Error())
		return nil, err
	}

	return config, nil
}
