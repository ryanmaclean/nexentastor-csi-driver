package config_test

import (
	"testing"

	"github.com/Nexenta/nexentastor-csi-driver/src/config"
)

var testConfigParams = map[string]string{
	"Address":        "https://10.1.1.1:8443,https://10.1.1.2:8443",
	"Username":       "usr",
	"Password":       "pwd",
	"DefaultDataset": "poolA/datasetA",
	"DefaultDataIp":  "20.1.1.1",
}

func testParam(t *testing.T, name, expected, given string) {
	if expected != given {
		t.Errorf("Param '%v' expected to be '%v', but got '%v' instead", name, expected, given)
	}
}

func TestConfig_Full(t *testing.T) {
	path := "./_fixtures/test-config-full"

	c, err := config.New(path)
	if err != nil {
		t.Fatalf("cannot read config file '%v': %v", path, err)
	}

	testParam(t, "Address", testConfigParams["Address"], c.Address)
	testParam(t, "Username", testConfigParams["Username"], c.Username)
	testParam(t, "Password", testConfigParams["Password"], c.Password)
	testParam(t, "DefaultDataset", testConfigParams["DefaultDataset"], c.DefaultDataset)
	testParam(t, "DefaultDataIp", testConfigParams["DefaultDataIp"], c.DefaultDataIP)
}

func TestConfig_Short(t *testing.T) {
	path := "./_fixtures/test-config-short"

	c, err := config.New(path)
	if err != nil {
		t.Fatalf("cannot read config file '%v': %v", path, err)
	}

	testParam(t, "Address", testConfigParams["Address"], c.Address)
	testParam(t, "Username", testConfigParams["Username"], c.Username)
	testParam(t, "Password", testConfigParams["Password"], c.Password)
	testParam(t, "DefaultDataset", "", c.DefaultDataset)
	testParam(t, "DefaultDataIp", "", c.DefaultDataIP)
}

func TestConfig_Not_Valid(t *testing.T) {

	t.Run("should return an error if config file if not valid", func(t *testing.T) {
		path := "./_fixtures/test-config-not-valid"
		c, err := config.New(path)
		if err == nil {
			t.Fatalf("not valid '%v' config file should return an error, but got this: %v", path, c)
		}
	})

	t.Run("should return nan error if file not exists", func(t *testing.T) {
		path := "./_fixtures/dir-without-config"
		c, err := config.New(path)
		if err == nil {
			t.Fatalf("not existing config file '%v' returns config: %v", path, c)
		}
	})
}
