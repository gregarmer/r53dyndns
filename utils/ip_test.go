package utils

import "testing"

func TestStub(t *testing.T) {
	if "foo" != "foo" {
		t.Fatal("something is really wrong")
	}
}

func TestValidIPValid(t *testing.T) {
	if !ValidIP("127.0.0.1") {
		t.Fatal("127.0.0.1 is a valid IP")
	}
}

func TestValidIPInvalid(t *testing.T) {
	if ValidIP("foo") {
		t.Fatal("foo is not a valid IP")
	}
}
