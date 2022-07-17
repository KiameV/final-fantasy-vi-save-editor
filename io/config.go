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
	WindowX       int    `json:"width"`
	WindowY       int    `json:"height""`
	SaveDir       string `json:"dir"`
	AutoEnableCmd bool   `json:"autoEnableCmd"`
}

func GetConfig() *ConfigData {
	return &config
}

func init() {
	if b, err := os.ReadFile(filepath.Join(global.PWD, file)); err == nil {
		_ = json.Unmarshal(b, &config)
	}
}

func SaveConfig() {
	if f, e1 := os.Create(filepath.Join(global.PWD, "ff6editor.config")); e1 == nil {
		if config.WindowX == 0 {
			config.WindowX = global.WindowWidth
		}
		if config.WindowY == 0 {
			config.WindowY = global.WindowHeight
		}
		b, err := json.Marshal(&config)
		if err == nil {
			os.WriteFile(filepath.Join(global.PWD, file), b, 0755)
		}
		_, _ = f.Write(b)
	}
}
