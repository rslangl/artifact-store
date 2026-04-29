package config

import (
	"testing"
)

func TestCreate(t *testing.T) {
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
				Storage: StorageConfig{
					Backend: "fs",
					Fs: FsConfig{
						Path: "/tmp/artifacts/",
					},
				},
			},
			false,
		},
		// Test 2: path provided, read contents
		{
			"config_test.yaml",
			Config{
				Storage: StorageConfig{
					Backend: "fs",
					Fs: FsConfig{
						Path: "/tmp/artifacts/",
					},
				},
			},
			false,
		},
	}

	for _, testInput := range tests {
		t.Run(testInput.path, func(t *testing.T) {
			cfg, err := New(testInput.path)
			if err != nil {
				t.Errorf("%v", err)
			}

			if (err != nil) != testInput.wantErr {
				t.Errorf("expected error: %v, got error: %v", testInput.wantErr, err)
			}

			if cfg.Storage.Backend != testInput.expected.Storage.Backend {
				t.Errorf("expected %v, got %v", testInput.expected.Storage.Backend, cfg.Storage.Backend)
			}
		})
	}
}
