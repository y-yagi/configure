package configure

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

// ConfigDir return config directory
func ConfigDir(name string) string {
	var dir string

	if runtime.GOOS == "windows" {
		dir = os.Getenv("APPDATA")
		if dir == "" {
			dir = filepath.Join(os.Getenv("USERPROFILE"), "Application Data", name)
		}
		dir = filepath.Join(dir, name)
	} else {
		dir = filepath.Join(os.Getenv("HOME"), ".config", name)
	}
	return dir
}

// Load loads config file and set result to cfg
func Load(name string, cfg interface{}) error {
	dir := ConfigDir(name)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create directory: %v", err)
	}
	file := filepath.Join(dir, "config.toml")

	_, err := os.Stat(file)
	if err == nil {
		_, err := toml.DecodeFile(file, cfg)
		if err != nil {
			return err
		}
		return nil
	}

	if !os.IsNotExist(err) {
		return err
	}

	return nil
}
