package dyndns

import "testing"

func TestStub(t *testing.T) {
	if "foo" != "foo" {
		t.Fatal("something is really wrong")
	}
}
