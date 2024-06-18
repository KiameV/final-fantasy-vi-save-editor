package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"pixel-remastered-save-editor/global"
)

const file = "prSaveEditor.config"

var config ConfigData

type ConfigData struct {
	SaveDir1 string `json:"dir1"`
	SaveDir2 string `json:"dir2"`
	SaveDir3 string `json:"dir3"`
	SaveDir4 string `json:"dir4"`
	SaveDir5 string `json:"dir5"`
	SaveDir6 string `json:"dir6"`
}

func init() {
	if b, err := os.ReadFile(filepath.Join(global.PWD, file)); err == nil {
		_ = json.Unmarshal(b, &config)
	}
}

func saveConfig() {
	if f, e1 := os.Create(filepath.Join(global.PWD, file)); e1 == nil {
		b, err := json.Marshal(&config)
		if err == nil {
			_ = os.WriteFile(filepath.Join(global.PWD, file), b, 755)
		}
		_, _ = f.Write(b)
	}
}

func Dir(saveType global.Game) (dir string) {
	switch saveType {
	case global.One:
		dir = config.SaveDir1
	case global.Two:
		dir = config.SaveDir2
	case global.Three:
		dir = config.SaveDir3
	case global.Four:
		dir = config.SaveDir4
	case global.Five:
		dir = config.SaveDir5
	case global.Six:
		dir = config.SaveDir6
	}
	if dir == "" {
		dir = "."
	}
	return
}

func SetSaveDir(saveType global.Game, dir string) {
	switch saveType {
	case global.One:
		config.SaveDir1 = dir
	case global.Two:
		config.SaveDir2 = dir
	case global.Three:
		config.SaveDir3 = dir
	case global.Four:
		config.SaveDir4 = dir
	case global.Five:
		config.SaveDir5 = dir
	default: // global.Six
		config.SaveDir6 = dir
	}
	saveConfig()
}
