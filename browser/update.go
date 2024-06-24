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
	tagUrl = `https://api.github.com/repos/KiameV/final-fantasy-pr-save-editor/tags`
	relUrl = `https://github.com/KiameV/final-fantasy-pr-save-editor/releases/%s`

	Version = "0.5.0"
)

type (
	tag struct {
		Name string `json:"name"`
	}
	comp struct {
		version [3]int
	}
)

func CheckForUpdate() (hasNewer bool, version string, err error) {
	var (
		r          *http.Response
		b          []byte
		tags       []tag
		current, _ = newComparable(Version)
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
			if other, ok := newComparable(t.Name); ok {
				if IsNewer(current, other) {
					current = other
					hasNewer = true
					version = t.Name
				}
			}
		}
	}
	return
}

func newComparable(v string) (c comp, found bool) {
	if strings.Contains(v, ".") && !strings.Contains(v, "_") {
		var (
			sb  strings.Builder
			err error
		)
		for _, r := range v {
			if r == '.' || (r >= '0' && r <= '9') {
				sb.WriteRune(r)
			}
		}
		for i, j := range strings.Split(sb.String(), ".") {
			if c.version[i], err = strconv.Atoi(j); err != nil {
				found = false
				return
			}
		}
		found = true
	}
	return
}

func IsNewer(current, other comp) bool {
	for i := 0; i < 3; i++ {
		c := current.version[i]
		o := other.version[i]
		if c < o {
			return true
		} else if c != o {
			return false
		}
	}
	return false
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
