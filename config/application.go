package config

// ApplicationConfig holds data necessary for App configuration
type ApplicationConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

// Init initializes ApplicationConfig
func (cfg ApplicationConfig) Init() error {
	return nil
}

type parseableApplicationConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

func (p parseableApplicationConfig) normalize() ApplicationConfig {
	return ApplicationConfig{
		Name:        p.Name,
		Description: p.Description,
		Version:     p.Version,
	}
}
