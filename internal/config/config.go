package config

import(
	"fmt"
)

type ConfigStorage struct {
	Enabled []string
}

type Config struct {
	StorageConfig ConfigStorage
}

func (cfg *Config) Create(path string) error {
	if path != "" {
		// TODO: read config file from disk
	} else {
		// TODO: read default config
	}
	return nil
}

func (cfg *Config) ToString() string {
	output := fmt.Sprintf("storages : %s", cfg.StorageConfig.Enabled)
	return output
}
