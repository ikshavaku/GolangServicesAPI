package utils

import (
	"fmt"
	"net/url"
	"sync"

	"github.com/joeshaw/envdecode"
)

type ConfServer struct {
	Port int `env:"SERVER_PORT,required"`
}

type PostgresConfig struct {
	PostgresPassword string `env:"db_password" validate:"required"`
	PostgresDBName   string `env:"db_name" validate:"required"`
	PostgresPort     string `env:"db_port" validate:"required"`
	PostgresUser     string `env:"db_user" validate:"required"`
	PostgresServer   string `env:"db_server" validate:"required"`
	PostgresDBSchema string `env:"db_schema" validate:"required"`
}

type GlobalConfig struct {
	Env      string `env:"env" validate:"required"`
	Server   ConfServer
	Postgres PostgresConfig `env:"postgres" validate:"required"`
}

var globalConfig GlobalConfig
var configOnce sync.Once

func InitConfig() error {
	var err error
	configOnce.Do(func() {
		err = envdecode.StrictDecode(&globalConfig)
	})
	return err
}

func GetConfig() GlobalConfig {
	return globalConfig
}

func (config PostgresConfig) PostgresConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=%s", url.QueryEscape(config.PostgresUser), url.QueryEscape(config.PostgresPassword), config.PostgresServer, config.PostgresPort, config.PostgresDBName, config.PostgresDBSchema)
}
