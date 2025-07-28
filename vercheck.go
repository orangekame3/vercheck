// Package vercheck provides a lightweight library for checking newer versions
// of CLI tools by querying GitHub releases and notifying users about available updates.
package vercheck

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/orangekame3/vercheck/internal"
)

// Options contains configuration for version checking.
type Options struct {
	CurrentVersion string // Required: Current CLI version (e.g., "v1.2.3")
	RepoOwner      string // Required: GitHub repository owner
	RepoName       string // Required: GitHub repository name
	UpdateCommand  string // Optional: Custom update command (auto-detected if empty)
	// Silent suppresses error logging (useful for optional version checks)
	Silent bool
}

// Check performs a version check against GitHub releases with a default 10-second timeout.
// It compares the current version with the latest release and displays update notifications if needed.
func Check(options Options) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	CheckWithContext(ctx, options)
}

// CheckWithContext performs a version check with a custom context for timeout and cancellation control.
// It compares the current version with the latest release and displays update notifications if needed.
func CheckWithContext(ctx context.Context, options Options) {
	if options.CurrentVersion == "" || options.RepoOwner == "" || options.RepoName == "" {
		return
	}

	latestVersion, err := getLatestReleaseWithContext(ctx, options.RepoOwner, options.RepoName)
	if err != nil {
		// Log error only if not in silent mode
		// This allows the CLI tool to continue functioning even if version check fails
		if !options.Silent {
			log.Printf("vercheck: failed to get latest release: %v", err)
		}
		return
	}

	if !isNewer(latestVersion, options.CurrentVersion) {
		return
	}

	updateCmd := options.UpdateCommand
	if updateCmd == "" {
		execPath, err := os.Executable()
		if err == nil {
			source := detectInstallSource(execPath)
			updateCmd = getDefaultUpdateCommand(source, options.RepoOwner, options.RepoName)
		}
	}

	fmt.Printf("🔔 New version %s is available! You're using %s.\n", latestVersion, options.CurrentVersion)
	if updateCmd != "" {
		fmt.Printf("👉 Update with: %s\n", updateCmd)
	}
}

func isNewer(latest, current string) bool {
	return internal.CompareVersions(latest, current) > 0
}
