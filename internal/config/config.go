package config

import(
	"os"
	//"log"
	"fmt"

	"go.yaml.in/yaml/v4"
)

type ConfigStorage struct {
	Enabled []string `yaml:,enabled`
}

type Config struct {
	Storage ConfigStorage `yaml:storage`
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
		cfg.Storage = ConfigStorage{Enabled: []string{"fs"}}
	}
	return nil
}

func (cfg *Config) ToString() string {
	output := fmt.Sprintf("storages : %s", cfg.Storage.Enabled)
	return output
}
