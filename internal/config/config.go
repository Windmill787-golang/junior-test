package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

type Server struct {
	MaxHeaderMB  int
	ReadTimeOut  int
	WriteTimeOut int
	Port         string
}

type Auth struct {
	AccessTokenTTL  string
	RefreshTokenTTL string
}

type Config struct {
	Env      string
	Postgres Postgres
	Server   Server
	Auth     Auth
}

func New(folder, filename string) (*Config, error) {
	cfg := new(Config)

	//set default app environment
	viper.SetDefault("env", "local")

	//load from directory
	viper.SetConfigName(filename)
	viper.AddConfigPath(folder)

	//load env
	viper.AutomaticEnv()

	//bind env variables
	//TODO: think of a better solution for this
	if err := viper.BindEnv("env", "APP_ENV"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("server.port", "SERVER_PORT"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("postgres.host", "POSTGRES_HOST"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("postgres.port", "POSTGRES_PORT"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("postgres.username", "POSTGRES_USERNAME"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("postgres.password", "POSTGRES_PASSWORD"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("postgres.database", "POSTGRES_DATABASE"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("postgres.sslmode", "POSTGRES_SSLMODE"); err != nil {
		return nil, err
	}

	//read
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return nil, fmt.Errorf("config file not found")
		} else {
			return nil, err
		}
	}

	//get config of app environment
	env := viper.Get("env")
	viper.SetConfigName(env.(string))

	//merge with default config
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	//unmarshal
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
