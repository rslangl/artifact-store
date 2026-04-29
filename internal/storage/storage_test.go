package storage

import (
	"testing"
	"os"

	"artifacts/internal/config"
	"artifacts/internal/storage/backend"
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
					Path: "/tmp/artifacts/",
				},
			},
			backend.FileSystem{
				Path: os.DirFS("/tmp/artifacts/"),
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
