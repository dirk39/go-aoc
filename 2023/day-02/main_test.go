package main

import "testing"

func TestSolution1(t *testing.T) {
	want := 8
	got1, _ := Solve("input-step1.txt")
	if got1 != want {
		t.Errorf("Solve(input-step1.txt) = %d; want %d", got1, want)
	}
}

func TestSolution2(t *testing.T) {
	want := 2286
	_, got2 := Solve("input-step1.txt")
	if got2 != want {
		t.Errorf("Solve(input-step1.txt) = %d; want %d", got2, want)
	}
}
