package configure_test

import (
	"testing"

	"github.com/y-yagi/configure"
)

func TestConfigDir(t *testing.T) {
	configDir := configure.ConfigDir("myrepo")
	expect := "/home/yaginuma/.config/myrepo"

	if configDir != expect {
		t.Errorf("Expect condig dir is %s, but %s", expect, configDir)
	}
}
