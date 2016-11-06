package dyndns

import "testing"

func TestStub(t *testing.T) {
	if "foo" != "foo" {
		t.Fatal("something is really wrong")
	}
}

func TestValidIPValid(t *testing.T) {
	dyndns := Dyndns{}
	if !dyndns.ValidIP("127.0.0.1") {
		t.Fatal("127.0.0.1 is a valid IP")
	}
}

func TestValidIPInvalid(t *testing.T) {
	dyndns := Dyndns{}
	if dyndns.ValidIP("foo") {
		t.Fatal("foo is not a valid IP")
	}
}
