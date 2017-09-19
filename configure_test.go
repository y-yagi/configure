package configure_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/y-yagi/configure"
)

func TestConfigDir(t *testing.T) {
	configDir := configure.ConfigDir("myrepo")
	expect := filepath.Join(os.Getenv("HOME"), ".config", "myrepo")

	if configDir != expect {
		t.Errorf("Expect condig dir is %s, but %s", expect, configDir)
	}
}
