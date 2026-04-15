package config

import(
	"os"
	"log"
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
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error occurred while reading config file '%v': %v", path, err)
		}
		log.Println("%v", data)
	} else {
		cfg.StorageConfig = ConfigStorage{Enabled: []string{"fs","nas","s3"}}
	}
	return nil
}

func (cfg *Config) ToString() string {
	output := fmt.Sprintf("storages : %s", cfg.StorageConfig.Enabled)
	return output
}
