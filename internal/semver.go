// Package internal provides internal utilities for semantic version comparison.
package internal

import (
	"strconv"
	"strings"
)

// CompareVersions compares two semantic versions
// Returns:
//
//	-1 if v1 < v2
//	 0 if v1 == v2
//	 1 if v1 > v2
func CompareVersions(v1, v2 string) int {
	v1 = normalizeVersion(v1)
	v2 = normalizeVersion(v2)

	v1Base, v1Pre := splitPreRelease(v1)
	v2Base, v2Pre := splitPreRelease(v2)

	if cmp := compareBasicVersions(v1Base, v2Base); cmp != 0 {
		return cmp
	}

	return comparePreReleases(v1Pre, v2Pre)
}

func normalizeVersion(version string) string {
	return strings.TrimPrefix(strings.TrimPrefix(version, "v"), "V")
}

func compareBasicVersions(v1, v2 string) int {
	parts1 := parseVersion(v1)
	parts2 := parseVersion(v2)

	for i := 0; i < 3; i++ {
		if parts1[i] > parts2[i] {
			return 1
		}
		if parts1[i] < parts2[i] {
			return -1
		}
	}
	return 0
}

func comparePreReleases(v1Pre, v2Pre string) int {
	// No pre-release > pre-release (1.0.0 > 1.0.0-alpha)
	if v1Pre == "" && v2Pre != "" {
		return 1
	}
	if v1Pre != "" && v2Pre == "" {
		return -1
	}
	if v1Pre != "" && v2Pre != "" {
		return strings.Compare(v1Pre, v2Pre)
	}
	return 0
}

func splitPreRelease(version string) (base, preRelease string) {
	// Remove build metadata first (anything after +)
	if idx := strings.Index(version, "+"); idx != -1 {
		version = version[:idx]
	}

	// Split pre-release (anything after -)
	if idx := strings.Index(version, "-"); idx != -1 {
		return version[:idx], version[idx+1:]
	}
	return version, ""
}

func parseVersion(v string) [3]int {
	var result [3]int

	parts := strings.Split(v, ".")
	for i := 0; i < 3 && i < len(parts); i++ {
		num, err := strconv.Atoi(parts[i])
		if err == nil {
			result[i] = num
		}
	}

	return result
}
