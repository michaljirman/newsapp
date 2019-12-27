package configs

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

// Config - main config struct
type Config struct {
	ServiceName string `env:"SVC_NAME"`
	DebugAddr   string `env:"DEBUG_ADDRESS"`
	GRPCAddr    string `env:"GRPC_ADDRESS"`
	Db          DBConfig
	URLBox      URLBoxConfig
}

type URLBoxConfig struct {
	URL    string `env:"URLBOX_API_URL"`
	Token  string `env:"URLBOX_API_TOKEN"`
	Secret string `env:"URLBOX_API_SECRET"`
}

// DBConfig - database configuration struct
type DBConfig struct {
	Host                    string `env:"DB_HOST"`
	Port                    int    `env:"DB_PORT"`
	User                    string `env:"DB_USER"`
	Password                string `env:"DB_PASSWORD"`
	Name                    string `env:"DB_NAME"`
	MaxOpenConns            int    `env:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns            int    `env:"DB_MAX_IDLE_CONNS"`
	ConnMaxLifetime         int    `env:"DB_CONN_MAX_LIFETIME"` // seconds
	LockTimeout             int    `env:"DB_LOCK_TIMEOUT"`      // ms
	MigrationsDirectoryPath string `env:"MIGRATIONS_DIRECTORY_PATH"`
}

func (c DBConfig) GetDSN() string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.Name)
	return dsn
}

// Get - Loads config from env
func Get() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	if err := env.Parse(&c.Db); err != nil {
		return nil, errors.Wrap(err, "failed to load Db config")
	}

	if err := env.Parse(&c.URLBox); err != nil {
		return nil, errors.Wrap(err, "failed to load URLBox config")
	}

	return c, nil
}
