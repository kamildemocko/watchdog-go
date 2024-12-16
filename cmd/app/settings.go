package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Settings struct {
	Processes      []string `toml:"processes"`
	LogFile        string   `toml:"log_file"`
	RefreshSeconds int      `toml:"refresh_seconds"`
}

func ParseSettings(path string) (Settings, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		return Settings{}, err
	}

	var settings Settings

	_, err = toml.Decode(string(buffer), &settings)
	if err != nil {
		return Settings{}, err
	}

	return settings, nil
}
