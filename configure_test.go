package configure

import (
	"os"
	"path/filepath"
	"testing"
)

type config struct {
	Path string
	Max  int
}

func TestConfigDir(t *testing.T) {
	configDir := ConfigDir("myrepo")
	expect := filepath.Join(os.Getenv("HOME"), ".config", "myrepo")

	if configDir != expect {
		t.Fatalf("Expect condig dir is %s, but %s", expect, configDir)
	}
}

func TestSave(t *testing.T) {
	name := "configure_test"
	if err := Save(name, config{Path: "dummy", Max: 20}); err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	configDir := ConfigDir(name)
	defer os.RemoveAll(configDir)

	var cfg config
	if err := Load(name, &cfg); err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if cfg.Path != "dummy" {
		t.Fatalf("Expect is %s, but %s", "dummy", cfg.Path)
	}

	if cfg.Max != 20 {
		t.Fatalf("Expect is %d, but %d", 20, cfg.Max)
	}
}
