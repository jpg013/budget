package config

import "fmt"

// PostgresConfig holds data necessary for PostgreSQL configuration
type PostgresConfig struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

// Init initializes postgres config default values
func (cfg PostgresConfig) Init() error {
	return nil
}

// MakeConnectionString formats the config values with password to a sql connection string
func (cfg PostgresConfig) MakeConnectionString(pass string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Username,
		pass,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}

type parseablePostgresConfig struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func (p parseablePostgresConfig) normalize() PostgresConfig {
	return PostgresConfig{
		Username: p.Username,
		Host:     p.Host,
		Port:     p.Port,
		Database: p.Database,
	}
}
