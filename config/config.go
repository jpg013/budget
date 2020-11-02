package config

// Configuration for the application
type Configuration struct {
	Postgres    PostgresConfig    `json:"postgres"`
	Application ApplicationConfig `json:"application"`
}

// parseableConfiguration represents the raw config values before being normalized
type parseableConfiguration struct {
	Postgres    parseablePostgresConfig    `json:"postgres"`
	Application parseableApplicationConfig `json:"application"`
}

func (cfg parseableConfiguration) normalize() Configuration {

	return Configuration{
		Postgres: cfg.Postgres.normalize(),
	}
}
