package file

import (
	"encoding/json"
	"os"
	"strings"
)

func printFile(name string, b []byte) {
	if os.Getenv("PR_PRINT") == "true" {
		s := strings.ReplaceAll(string(b), "\\", "")
		s = strings.ReplaceAll(s, "\"{", "{")
		s = strings.ReplaceAll(s, "}\"", "}")
		var m map[string]any
		if err := json.Unmarshal([]byte(s), &m); err == nil {
			if b, err = json.MarshalIndent(m, "", "\t"); err == nil {
				_ = os.WriteFile(name, b, 0755)
				return
			}
		}
		_ = os.WriteFile(name, []byte(s), 0755)
	}
}
