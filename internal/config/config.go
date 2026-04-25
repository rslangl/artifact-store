package config

import(
	"os"
	"fmt"

	"go.yaml.in/yaml/v4"
)

type FsConfig struct {
	Path string `yaml:path`
}

type StorageConfig struct {
	Backend string `yaml:backend`
	Fs FsConfig `yaml:fs,omitempty`
}

type ServiceConfig struct {
	Address string `yaml:address`
}

type Config struct {
	Storage StorageConfig `yaml:storage`
	Service ServiceConfig `yaml:service`
}

func New(path string) (Config, error) {
	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return Config{}, fmt.Errorf("error occurred while reading config file '%v': %v", path, err)
		} else {
			cfg := Config{}
			err := yaml.Unmarshal(data, &cfg)
			if err != nil {
				return Config{}, fmt.Errorf("error occurred while constructing config from file '%v': %v", path, err)
			}
		}
	}
	return Config{
		Service: ServiceConfig{
			Address: "0.0.0.0:8080",
		},
		Storage: StorageConfig{
			Backend: "fs",
			Fs: FsConfig{
				Path: "/tmp/artifacts",
			},
		},
	}, nil
}

func (cfg *Config) ToString() string {
	output := fmt.Sprintf("storages : %s", cfg.Storage.Backend)
	return output
}
