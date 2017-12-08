package main

import "testing"

func TestTimes(t *testing.T) {
	total := Times(5, 5)
	if total != 25 {
		t.Errorf("Perkalian Error, got: %d, want: %d.", total, 10)
	}
}
