package config

import (
	"testing"
)

func isEqual(v1, v2 []string) bool {
	if len(v1) != len(v2) {
		return false
	}
	for idx := range v1 {
		if v1[idx] != v2[idx] {
			return false
		}
	}
	return true
}

func TestCreate(t *testing.T) {

	cfg := &Config{}

	tests := []struct{
		// The argument passed to the function
		path string
		// The expected output
		expected Config
		// Whether we want this test to fail
		wantErr bool
	}{
		// Test 1: no path, read defaults
		{
			"",
			Config{
				Storage: ConfigStorage{
					Enabled: []string{"fs"},
				},
			},
			false,
		},
		// Test 2: path provided, read contents
		{
			"config_test.yaml",
			Config{
				Storage: ConfigStorage{
					Enabled: []string{"fs", "nas", "s3"},
				},
			},
			false,
		},
	}

	for _, testInput := range tests {
		t.Run(testInput.path, func(t *testing.T) {
			err := cfg.Create(testInput.path)

			if (err != nil) != testInput.wantErr {
				t.Errorf("expected error: %v, got error: %v", testInput.wantErr, err)
			}

			if !isEqual(cfg.Storage.Enabled, testInput.expected.Storage.Enabled) {
				t.Errorf("expected %v, got %v", testInput.expected.Storage.Enabled, cfg.Storage.Enabled)
			}
		})
	}
}
