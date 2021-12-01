package main

import "testing"

const ANSWER = 1523

func TestSolution(t *testing.T) {
	s := solution()
	if s != ANSWER {
		t.Errorf("Expected: %d, got %d", ANSWER, s)
	}
}
