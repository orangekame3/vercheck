package vercheck

import (
	"fmt"
	"strings"
)

func detectInstallSource(execPath string) string {
	switch {
	case strings.Contains(execPath, "/Cellar/"):
		return "homebrew"
	default:
		return "go-install"
	}
}

func getDefaultUpdateCommand(source, owner, repo string) string {
	switch source {
	case "homebrew":
		return fmt.Sprintf("brew upgrade %s", repo)
	case "go-install":
		return fmt.Sprintf("go install github.com/%s/%s@latest", owner, repo)
	default:
		return ""
	}
}
