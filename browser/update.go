package browser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

const (
	tagUrl = `https://api.github.com/repos/KiameV/final-fantasy-vi-save-editor/tags`
	relUrl = `https://github.com/KiameV/final-fantasy-vi-save-editor/releases/%s`
)

type tag struct {
	Name string `json:"name"`
}

func CheckForUpdate(current string) (hasNewer bool, version string, err error) {
	var (
		r    *http.Response
		b    []byte
		tags []tag
		vn   = versionToInt(current)
	)
	if r, err = http.Get(tagUrl); err != nil {
		return
	}
	if b, err = io.ReadAll(r.Body); err != nil {
		return
	}
	if err = json.Unmarshal(b, &tags); err != nil {
		return
	}

	for _, t := range tags {
		if strings.Contains(t.Name, ".") {
			i := versionToInt(t.Name)
			if i > vn {
				hasNewer = true
				version = t.Name
				vn = i
			}
		}
	}
	return
}

func versionToInt(v string) int {
	if i, err := strconv.ParseInt(strings.ReplaceAll(v, ".", ""), 10, 32); err != nil {
		return 0
	} else {
		return int(i)
	}
}

func Update(tag string) (err error) {
	url := fmt.Sprintf(relUrl, tag)
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return
}
