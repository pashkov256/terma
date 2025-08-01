package config

import (
	"encoding/json"
	"os"

	"github.com/pashkov256/terma/internal/path"
)

type Config struct {
	APIKey    string   `json:"api_key"`
	Language  string   `json:"language"`
	Units     string   `json:"units"` // "metric", "imperial"
	Theme     string   `json:"theme"`
	Favorites []string `json:"favorites"`
}

func (c *Config) SaveConfig() error {
	configJson, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(os.Getenv("HOME")+"/"+path.AppDir+"/"+path.ConfigFile, configJson, 0644)
}

func (c *Config) GetConfig(path string) (*Config, error) {
	configJson, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configJson, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
