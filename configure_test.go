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
		t.Fatalf("Expect config dir is %s, but %s", expect, configDir)
	}
}

func TestCustomizeConfigDir(t *testing.T) {
	os.Setenv("CONFIGURE_DIRECTORY", "/tmp/config")
	defer os.Unsetenv("CONFIGURE_DIRECTORY")

	configDir := ConfigDir("myrepo")

	expect := filepath.Join("/tmp/config", "myrepo")

	if configDir != expect {
		t.Fatalf("Expect config dir is %s, but %s", expect, configDir)
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

func TestSaveWithBackup(t *testing.T) {
	name := "configure_test"
	if err := Save(name, config{Path: "dummy", Max: 20}); err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	configDir := ConfigDir(name)
	defer os.RemoveAll(configDir)

	tempDir, err := os.MkdirTemp("", "configuretest")
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	defer os.RemoveAll(tempDir)

	c := Configure{Name: name, BackupDir: tempDir}
	if err := c.Save(config{Path: "new_dummy", Max: 120}); err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	var cfg config
	if err := c.Load(&cfg); err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if cfg.Path != "new_dummy" {
		t.Fatalf("Expect is %s, but %s", "new_dummy", cfg.Path)
	}

	if cfg.Max != 120 {
		t.Fatalf("Expect is %d, but %d", 120, cfg.Max)
	}

	got, err := os.ReadFile(filepath.Join(tempDir, "config.toml"))
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	expected := `Max = 20
Path = "dummy"
`
	if string(got) != expected {
		t.Fatalf("Expect is %v, but %v", expected, string(got))
	}

}

func TestExist(t *testing.T) {
	name := "configure_test"

	if Exist(name) {
		t.Fatalf("Expect 'Exist' is return false, but true")
	}

	var cfg config
	Save(name, cfg)
	defer os.RemoveAll(ConfigDir(name))

	if !Exist(name) {
		t.Fatalf("Expect 'Exist' is return true, but false")
	}
}
