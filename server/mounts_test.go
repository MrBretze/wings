package server

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsDevicePath(t *testing.T) {
	if _, err := os.Stat("/dev/null"); err != nil {
		t.Skip("/dev/null is not available")
	}

	dir := t.TempDir()
	file := filepath.Join(dir, "file")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	cases := map[string]bool{
		"/dev/null":                   true,
		"/dev":                        true,
		dir:                           false,
		file:                          false,
		filepath.Join(dir, "missing"): false,
	}
	for path, want := range cases {
		if got := isDevicePath(path); got != want {
			t.Errorf("isDevicePath(%q) = %v, want %v", path, got, want)
		}
	}
}
