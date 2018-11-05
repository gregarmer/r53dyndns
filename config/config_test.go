package config

import (
	"fmt"
	"strings"
	"testing"
)

func TestPreFlight(t *testing.T) {
	conf := Config{}
	err := conf.PreFlight()
	if err == nil {
		t.Fatalf("error should be set")
	}
}

func TestCopy(t *testing.T) {
	conf := Config{AwsAccessKey: "foo"}
	c := conf.Copy()
	c.AwsAccessKey = "bar"
	if conf.AwsAccessKey == c.AwsAccessKey {
		t.Fatalf("config.Copy() should return a copy that doesn't affect the original")
	}
}

func TestConfigMissing(t *testing.T) {
	conf := Config{}
	err := conf.PreFlight()
	if !strings.Contains(fmt.Sprintf("%s", err), "missing AwsAccessKey") {
		t.Fatalf("PreFlight() should raise an error for missing AwsAccessKey")
	}

	conf = Config{AwsAccessKey: "foo"}
	err = conf.PreFlight()
	if !strings.Contains(fmt.Sprintf("%s", err), "missing AwsSecretKey") {
		t.Fatalf("PreFlight() should raise an error for missing AwsSecretKey")
	}

	conf = Config{AwsAccessKey: "foo", AwsSecretKey: "bar"}
	err = conf.PreFlight()
	if !strings.Contains(fmt.Sprintf("%s", err), "missing ZoneId") {
		t.Fatalf("PreFlight() should raise an error for missing ZoneId")
	}

	conf = Config{AwsAccessKey: "foo", AwsSecretKey: "bar", ZoneId: "baz"}
	err = conf.PreFlight()
	if err != nil {
		t.Fatalf("PreFlight() should not raise an error")
	}
}

func TestLoadConfig(t *testing.T) {
	conf := LoadConfig("./.r53dyndns")
	if conf.AwsAccessKey != "foo" {
		t.Fatalf("Config was not loaded correctly.")
	}
}

func TestSampleConfig(t *testing.T) {
	conf := SampleConfig()
	expected := `{
  "aws_access_key": "",
  "aws_secret_key": "",
  "zone_id": ""
}`

	if conf != expected {
		t.Fatalf("SampleConfig did not return expected config. "+
			"Got: \n\n%s \n\nbut expected \n\n%s", conf, expected)
	}
}
