package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	result := hello()
	if result != "Hello" {
		t.Error("unexpected result")
	}
}
