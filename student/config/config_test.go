package config

import (
	"testing"
)

func TestInit(t *testing.T) {
	t.Logf("should be empty: %#v", Database)
	Init("/Users/c/Desktop/StudentConf.yml")
	t.Logf("%#v", config)
	t.Logf("should not be empty: %#v", Database)
}
