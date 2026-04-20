package storage

import (
	"testing"
	"os"

	"artifact-store/internal/config"
	"artifact-store/internal/storage/backend"
)

func TestCreate(t *testing.T) {
	stg := &Storage{}

	tests := []struct{
		// The argument passed to the function
		config config.StorageConfig
		// The expected output
		expected Storage
		// Whether we want this test to fail
		wantErr bool
	}{
		// Test 1: simple file system backend only
		{
			config.StorageConfig{
				Enabled: []string{"fs"},
				Fs: config.FsConfig{
					Path: "/tmp/artifact-store/",
				},
			},
			Storage{
				FileSystem: backend.FileSystem{
					Path: os.DirFS("/tmp/artifact-store/"),
				},
			},
			false,
		},
	}

	for _, testInput := range tests {
		t.Run("", func(t *testing.T) {
			err := stg.Create(testInput.config)

			if (err != nil) != testInput.wantErr {
				t.Errorf("expected error: %v, got error: %v", testInput.wantErr, err)
			}

			if stg.FileSystem.Path != testInput.expected.FileSystem.Path {
				t.Errorf("output '%v' does not match the expected value: %v", stg.FileSystem.Path, testInput.expected.FileSystem.Path)
			}
		})
	}
}
