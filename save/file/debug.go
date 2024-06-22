package file

import (
	"os"
	"strings"
)

func printFile(name string, b []byte) {
	if os.Getenv("PR_PRINT") == "true" {
		s := strings.ReplaceAll(string(b), "\\", "")
		s = strings.ReplaceAll(s, "\"{", "\n{")
		s = strings.ReplaceAll(s, "}\"", "}")
		os.WriteFile(name, []byte(s), 0755)
	}
}
