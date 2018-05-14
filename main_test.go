package main

import (
	"testing"
)

func Testxxd(t *testing.T) {
	result := xxd("date")
	if result != "00000000: 6461 7465 0a                             date." {
		t.Errorf("Unexpected result : %s", result)
	}
}
