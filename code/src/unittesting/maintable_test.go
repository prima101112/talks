package main

import "testing"

func TestTimesTable(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{5, 5, 25},
		{6, 7, 42},
		{2, 9, 18},
	}

	for i, table := range tables {
		total := Times(table.x, table.y)
		if total != table.n {
			t.Errorf("Perkalian of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		} else {
			t.Logf("Test case %d berhasil", i)
		}
	}
}
