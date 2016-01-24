package config

import "testing"

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
