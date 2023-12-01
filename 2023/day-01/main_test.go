package main

import (
	"testing"
)

func TestBaseInput(t *testing.T) {
	want := 142
	got := Solve("test-input.txt")

	if got != want {
		t.Errorf("Solve(test-input.txt) = %d; want %d", got, want)
	}
}

func TestAdvancedInput(t *testing.T) {
	want := 281
	got := Solve("test2-input.txt")

	if got != want {
		t.Errorf("Solve(test2-input.txt) = %d; want %d", got, want)
	}
}
