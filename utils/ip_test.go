package utils

import (
	"fmt"
	"strings"
	"testing"
)

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

func TestGetExternalIP(t *testing.T) {
	ip, err := GetExternalIP()
	if err != nil {
		t.Fatal(err)
	}

	if !ValidIP(ip) {
		t.Fatalf("%s is not a valid IP address", ip)
	}
}

func TestGetInterfaceIP(t *testing.T) {
	ip, err := GetInterfaceIP("lo")
	if err != nil {
		t.Fatal(err)
	}

	if ip != "127.0.0.1" {
		t.Fatalf("%s != 127.0.0.1", ip)
	}
}

func TestGetInterfaceIPInvalidInterface(t *testing.T) {
	_, err := GetInterfaceIP("blahblahblah")
	if err == nil {
		t.Fatal("Expected error to occur")
	}

	if !strings.Contains(fmt.Sprintf("%s", err), "no such network interface") {
		t.Fatalf("Unexpected error: %s", err)
	}
}
