// Package main demonstrates usage of the vercheck library.
package main

import (
	"github.com/orangekame3/vercheck"
)

func main() {
	vercheck.Check(vercheck.Options{
		CurrentVersion: "v0.4.0",
		RepoOwner:      "orangekame3",
		RepoName:       "lazypoetry",
	})
}
