package configure

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
)

// ConfigDir return config directory
func ConfigDir(name string) string {
	var dir string

	dir = os.Getenv("CONFIGURE_DIRECTORY")
	if dir != "" {
		return filepath.Join(dir, name)
	}

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

// Save saves config to file
func Save(name string, cfg interface{}) error {
	dir := ConfigDir(name)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create directory: %v", err)
	}
	file := filepath.Join(dir, "config.toml")

	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("cannot create config file: %v", err)
	}
	return toml.NewEncoder(f).Encode(cfg)
}

// Edit run editor for edit config file
func Edit(name string, editor string) error {
	if len(editor) == 0 {
		return errors.New("editor is empty")
	}

	dir := ConfigDir(name)
	file := filepath.Join(dir, "config.toml")
	cmd := exec.Command(editor, file)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// Exist check config file exist or not
func Exist(name string) bool {
	dir := ConfigDir(name)
	filename := filepath.Join(dir, "config.toml")
	_, err := os.Stat(filename)
	return err == nil
}
