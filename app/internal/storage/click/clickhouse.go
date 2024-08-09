package click

import (
	"context"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Config struct {
	Hosts    []string `yaml:"hosts"`
	Database string   `yaml:"db_name"`
	Port     int      `yaml:"port" env-default:"9000"`
	Debug    bool     `yaml:"debug" env-default:"false"`
}

type Storage struct {
	db *driver.Conn
}

func New(cfg Config) (*Storage, error) {
	const op = "storage.clickhouse.New"

	ctx := context.Background()

	db, err := clickhouse.Open(&clickhouse.Options{
		Auth: clickhouse.Auth{
			Database: cfg.Database,
		},
		Addr:  cfg.GetDSN(),
		Debug: cfg.Debug,
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{{Name: "SMSC", Version: "0.0.1"}},
		},
		ConnOpenStrategy: clickhouse.ConnOpenRoundRobin,
	})

	if err := db.Ping(ctx); err != nil {
		var exception *clickhouse.Exception
		if errors.As(err, &exception) {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: &db}, err
}

func (c *Config) GetDSN() []string {

	var dsn []string

	for _, host := range c.Hosts {
		dsn = append(dsn, fmt.Sprintf("%s:%d", host, c.Port))
	}

	return dsn
}
