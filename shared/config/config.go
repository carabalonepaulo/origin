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

func LoadDefault[T any]() *T {
	return Load[T](PathOrDefault("./config.json"))
}

func Load[T any](path string) *T {
	config, err := TryLoad[T](path)
	if err != nil {
		panic(err)
	}
	return config
}

func TryLoad[T any](path string) (*T, error) {
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
