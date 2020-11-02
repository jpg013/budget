package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

// Parser interface declaration
type Parser interface {
	Parse(configFile string) (Configuration, error)
}

// NewParser factory returns a new parser
func NewParser() Parser {
	return parser{}
}

type parser struct {
}

// Parse implements the Parser interface
func (p parser) Parse(filePath string) (cfg Configuration, err error) {
	cfgPath := flag.String("p", filePath, "Path to the configuration filename")
	flag.Parse()
	data, err := ioutil.ReadFile(*cfgPath)
	if err != nil {
		return cfg, fmt.Errorf("Error parsing config file: %s", err.Error())
	}
	parseableCfg := new(parseableConfiguration)
	if err = json.Unmarshal(data, &parseableCfg); err != nil {
		return cfg, fmt.Errorf("Error parsing config file: %s", err.Error())
	}

	cfg = parseableCfg.normalize()

	// Init default values
	err = cfg.Application.Init()
	if err != nil {
		return cfg, fmt.Errorf("Error initializing application config: %s", err.Error())
	}

	err = cfg.Postgres.Init()
	if err != nil {
		return cfg, fmt.Errorf("Error initializing mysql config: %s", err.Error())
	}

	return cfg, err
}
