package config

import (
	"context"
	"errors"
	"fmt"
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
	once   sync.Once //nolint:gochecknoglobals // no need
	config *Config   //nolint:gochecknoglobals // no need
)

// getGoModRoot returns the absolute path to the root directory containing go.mod.
func getGoModRoot() (string, error) {
	output, err := exec.CommandContext(context.Background(), "go", "list", "-m", "-f", "{{.Dir}}").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func loadConfig(path string) (*viper.Viper, error) {
	v := viper.New()

	configRAW, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't load config file %s: %w", path, err)
	}

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if err = v.ReadConfig(configRAW); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError

		if errors.As(err, &configFileNotFoundError) {
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
		return nil, fmt.Errorf("unable to decode into struct, %w", err)
	}

	err = validator.GetInstance().Struct(&c)
	if err != nil {
		return nil, fmt.Errorf("can't validate config: %w", err)
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
					panic(fmt.Errorf("error loading config file: %w", err))
				}

				cfg, err := parseConfig(viperCfg)
				if err != nil {
					panic(fmt.Errorf("error parsing config file: %w", err))
				}

				config = cfg
			},
		)
	}

	return config
}
