package config

import(
	"os"
	//"log"
	"fmt"

	"go.yaml.in/yaml/v4"
)

type FsConfig struct {
	Path string `yaml:path`
}

type StorageConfig struct {
	Enabled []string `yaml:,enabled`
	Fs FsConfig `yaml:fs`
}

type ServiceConfig struct {
	Address string `yaml:address`
}

type Config struct {
	Storage StorageConfig `yaml:storage`
	Service ServiceConfig `yaml:service`
}

func (cfg *Config) Create(path string) error {
	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error occurred while reading config file '%v': %v", path, err)
		} else {
			err := yaml.Unmarshal(data, &cfg)
			if err != nil {
				return fmt.Errorf("error occurred while constructing config from file '%v': %v", path, err)
			}
		}
		//log.Println("%+v\n", data)
	} else {
		cfg.Storage = StorageConfig{Enabled: []string{"fs"}}
		cfg.Service = ServiceConfig{Address: "0.0.0.0:8080"}
	}
	return nil
}

func (cfg *Config) ToString() string {
	output := fmt.Sprintf("storages : %s", cfg.Storage.Enabled)
	return output
}
