package main

import "testing"

func TestAddition(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("Expected 2 + 3 to equal 5, but got %d", result)
	}
}

func TestSubtraction(t *testing.T) {
	result := subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 5 - 3 to equal 2, but got %d", result)
	}
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
