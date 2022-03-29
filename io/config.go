package io

import (
	"encoding/json"
	"ffvi_editor/global"
	"os"
	"path/filepath"
)

const file = "ff6editor.config"

var config ConfigData

type ConfigData struct {
	WindowX int    `json:"width"`
	WindowY int    `json:"height""`
	SaveDir string `json:"dir"`
}

func GetConfig() *ConfigData {
	return &config
}

func init() {
	if b, err := os.ReadFile(filepath.Join(global.PWD, file)); err == nil {
		if err = json.Unmarshal(b, &config); err != nil {
			config.SaveDir = string(b)
			if config.SaveDir != "" {
				SaveConfig()
			}
		}
	}
}

func SaveConfig() {
	if f, e1 := os.Create(filepath.Join(global.PWD, "ff6editor.config")); e1 == nil {
		b, err := json.Marshal(&config)
		if err == nil {
			os.WriteFile(filepath.Join(global.PWD, file), b, 644)
		}
		_, _ = f.Write(b)
	}
}
