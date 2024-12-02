package config

import (
	"encoding/json"
	"os"

	"github.com/carabalonepaulo/origin/shared/services/scheduler"
)

type (
	Listener struct {
		Port         uint16 `json:"port"`
		MaxClients   int    `json:"max_clients"`
		InLimit      int    `json:"in_limit"`
		OutLimit     int    `json:"out_limit"`
		Channels     uint64 `json:"channels"`
		TickInterval string `json:"tick_interval"`
	}

	Database struct {
		Path         string `json:"path"`
		TickInterval string `json:"tick_interval"`
	}

	Config struct {
		Scheduler scheduler.Config `json:"scheduler"`
		Listener  `json:"listener"`
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
