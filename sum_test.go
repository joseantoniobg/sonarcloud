package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2,3)

	if result != 5 {
		t.Error("Result must be 5")
	}
}

func TestSub(t *testing.T) {
	result := sub(2,3)

	if result != -1 {
		t.Error("Result must be -1")
	}
}

func TestMult(t *testing.T) {
	result := mlt(2,3)

	if result != 6 {
		t.Error("Result must be 6")
	}
}

func TestDiv(t *testing.T) {
	result := div(6,2)

	if result != 3 {
		t.Error("Result must be 3")
	}
}