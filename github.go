package vercheck

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type githubRelease struct {
	TagName string `json:"tag_name"`
}

func getLatestReleaseWithContext(ctx context.Context, owner, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)

	// Create a transport with proper timeouts
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "vercheck/1.0 (+https://github.com/orangekame3/vercheck)")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("github API request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			// Log error but don't fail the operation
			// Response body close errors are typically not critical
			return
		}
	}()

	switch resp.StatusCode {
	case http.StatusOK:
		// Success, continue
	case http.StatusNotFound:
		return "", fmt.Errorf("repository or release not found")
	case http.StatusForbidden:
		// Check if it's rate limit
		if remaining := resp.Header.Get("X-RateLimit-Remaining"); remaining == "0" {
			resetTime := resp.Header.Get("X-RateLimit-Reset")
			return "", fmt.Errorf("github API rate limit exceeded, resets at %s", resetTime)
		}
		return "", fmt.Errorf("github API returned forbidden")
	default:
		return "", fmt.Errorf("github API returned status %d", resp.StatusCode)
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	return release.TagName, nil
}
