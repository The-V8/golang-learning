package main

import "testing"

func TestTestee(t *testing.T) {
	if !Testee() {
		t.Errorf("Failed to get boolean TRUE from test function.")
	}
}
