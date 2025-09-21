package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"yelp/internal/pkg/validator"

	"github.com/spf13/viper"
)

type Config struct {
	ExpirationWorker WorkerPool
}

var (
	once   sync.Once
	config *Config
)

// getGoModRoot returns the absolute path to the root directory containing go.mod.
func getGoModRoot() (string, error) {
	output, err := exec.Command("go", "list", "-m", "-f", "{{.Dir}}").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func loadConfig(path string) (*viper.Viper, error) {
	v := viper.New()

	configRAW, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't load config file %s: %s", path, err)
	}

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if err := v.ReadConfig(configRAW); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}

		return nil, err
	}

	return v, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	err = validator.GetInstance().Struct(&c)
	if err != nil {
		return nil, fmt.Errorf("can't validate config: %s", err.Error())
	}

	return &c, nil
}

func GetInstance() *Config {
	if config == nil {
		once.Do(
			func() {
				var err error

				path := "./internal/config/config.yaml"

				root, err := getGoModRoot()
				if err == nil {
					path = root + "/internal/config/config.yaml"
				}

				viperCfg, err := loadConfig(path)
				if err != nil {
					log.Fatalf("error loading config file: %s", err)
				}

				cfg, err := parseConfig(viperCfg)
				if err != nil {
					log.Fatalf("error parsing config file: %s", err)
				}

				config = cfg
			},
		)
	}

	return config
}
