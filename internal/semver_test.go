package internal

import "testing"

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		name     string
		v1       string
		v2       string
		expected int
	}{
		// Basic version comparisons
		{"equal versions", "1.0.0", "1.0.0", 0},
		{"v1 greater major", "2.0.0", "1.0.0", 1},
		{"v1 lesser major", "1.0.0", "2.0.0", -1},
		{"v1 greater minor", "1.2.0", "1.1.0", 1},
		{"v1 lesser minor", "1.1.0", "1.2.0", -1},
		{"v1 greater patch", "1.0.2", "1.0.1", 1},
		{"v1 lesser patch", "1.0.1", "1.0.2", -1},

		// With v prefix
		{"equal with v prefix", "v1.0.0", "v1.0.0", 0},
		{"mixed v prefix", "v1.0.0", "1.0.0", 0},
		{"capital V prefix", "V1.0.0", "v1.0.0", 0},

		// Pre-release versions
		{"stable > pre-release", "1.0.0", "1.0.0-alpha", 1},
		{"pre-release < stable", "1.0.0-alpha", "1.0.0", -1},
		{"pre-release comparison", "1.0.0-alpha", "1.0.0-beta", -1},
		{"pre-release comparison 2", "1.0.0-beta", "1.0.0-alpha", 1},
		{"pre-release with numbers", "1.0.0-alpha.1", "1.0.0-alpha.2", -1},

		// Build metadata (should be ignored)
		{"with build metadata", "1.0.0+build", "1.0.0", 0},
		{"both with build metadata", "1.0.0+build1", "1.0.0+build2", 0},
		{"pre-release with build", "1.0.0-alpha+build", "1.0.0-alpha", 0},

		// Real-world examples
		{"real world 1", "v0.0.9", "v0.0.10", -1},
		{"real world 2", "v1.2.3", "v1.3.0", -1},
		{"real world 3", "v2.0.0", "v1.9.9", 1},

		// Edge cases
		{"missing parts", "1", "1.0.0", 0},
		{"missing parts 2", "1.0", "1.0.0", 0},
		{"pre-release ordering", "1.0.0-1", "1.0.0-10", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CompareVersions(tt.v1, tt.v2)
			if result != tt.expected {
				t.Errorf("CompareVersions(%q, %q) = %d, want %d",
					tt.v1, tt.v2, result, tt.expected)
			}
		})
	}
}
