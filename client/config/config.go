package config

import (
	"encoding/json"
	"os"

	"github.com/carabalonepaulo/origin/shared/services/scheduler"
)

type (
	Config struct {
		Scheduler scheduler.Config `json:"scheduler"`
	}
)

func PathOrDefault(p string) (path string) {
	path = os.Getenv("CONFIG_PATH")
	if path == "" {
		path = p
	}
	return
}

func Load(path string) (*Config, error) {
	buff, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(buff, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
