package storage

import (
	"testing"
	"os"

	"artifact-store/internal/config"
	"artifact-store/internal/storage/backend"
)

func TestCreate(t *testing.T) {
	tests := []struct{
		// The argument passed to the function
		config config.StorageConfig
		// The expected output
		expected backend.FileSystem
		// Whether we want this test to fail
		wantErr bool
	}{
		// Test 1: simple file system backend only
		{
			config.StorageConfig{
				Backend: "fs",
				Fs: config.FsConfig{
					Path: "/tmp/artifact-store/",
				},
			},
			backend.FileSystem{
				Path: os.DirFS("/tmp/artifact-store/"),
			},
			false,
		},
	}

	for _, testInput := range tests {
		t.Run("", func(t *testing.T) {
			stg, err := New(testInput.config)

			if (err != nil) != testInput.wantErr {
				t.Errorf("expected error '%v', got error: %v", testInput.wantErr, err)
			}

			if _, ok := stg.(*backend.FileSystem); !ok {
				t.Errorf("expected type '%T', got type: %T", testInput.expected, stg)
			}
		})
	}
}
