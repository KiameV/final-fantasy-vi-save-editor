package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"ffvi_editor/global"
)

const (
	file = "ff6editor.config"
)

var (
	data d
)

type (
	d struct {
		WindowX       float32 `json:"width"`
		WindowY       float32 `json:"height""`
		SaveDir       string  `json:"dir"`
		AutoEnableCmd bool    `json:"autoEnableCmd"`
	}
)

func init() {
	if b, err := os.ReadFile(filepath.Join(global.PWD, file)); err == nil {
		_ = json.Unmarshal(b, &data)
	}
	if data.WindowX == 0 {
		data.WindowX = global.WindowWidth
	}
	if data.WindowY == 0 {
		data.WindowY = global.WindowHeight
	}
}

func WindowSize() (x, y float32) {
	x = data.WindowX
	y = data.WindowY
	return
}

func SaveDir() string {
	return data.SaveDir
}

func AutoEnableCmd() bool {
	return data.AutoEnableCmd
}

func SetWindowSize(x, y float32) {
	data.WindowX = x
	data.WindowY = y
	save()
}

func SetSaveDir(dir string) {
	data.SaveDir = dir
	save()
}

func SetAutoEnableCmd(v bool) {
	data.AutoEnableCmd = v
	save()
}

func save() {
	if f, e1 := os.Create(filepath.Join(global.PWD, "ff6editor.config")); e1 == nil {
		if data.WindowX == 0 {
			data.WindowX = global.WindowWidth
		}
		if data.WindowY == 0 {
			data.WindowY = global.WindowHeight
		}
		b, err := json.Marshal(&data)
		if err == nil {
			_ = os.WriteFile(filepath.Join(global.PWD, file), b, 0755)
		}
		_, _ = f.Write(b)
	}
}
