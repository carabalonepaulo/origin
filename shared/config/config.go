package config

import (
	"encoding/json"
	"os"
)

func PathOrDefault(p string) (path string) {
	path = os.Getenv("CONFIG_PATH")
	if path == "" {
		path = p
	}
	return
}

func Load[T any](path string) (*T, error) {
	buff, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config T
	err = json.Unmarshal(buff, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
