package main

import (
	"testing"
)

func TestDiablo(t *testing.T) {
	result := diablo(666)
	if result != 666 {
		t.Errorf("Main return was incorrect, got: %d, want: %d.", result, 666)
	}
}
