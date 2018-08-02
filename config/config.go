package config

import (
	"os"
	"time"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ShowInterval    int                  `yaml:"show_interval"`
	RefreshInterval int                  `yaml:"refresh_interval"`
	SocketPath      string               `yaml:"socker_path,omitempty"`
	Cues            map[string]CueConfig `yaml:"cues,flow"`
}

type CueConfig struct {
	CueType     string            `yaml:"type"`
	Default     bool              `yaml:"default,omitempty"`
	Description string            `yaml:"description"`
	CacheTTL    time.Duration     `yaml:"cache_ttl,omitempty"`
	Config      map[string]string `yaml:"config,omitempty"`
}

func NewConfig(path string) (*Config, error) {
	var cfg Config

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config file")
	}

	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to decode configuration")
	}

	return &cfg, nil
}
