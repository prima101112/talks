package main

import "testing"

func TestDevideTable(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{10, 5, 2},
		{30, 3, 10},
		{6, 2, 3},
	}

	for i, table := range tables {
		total := Devide(table.x, table.y)
		if total != table.n {
			t.Errorf("Pembagian of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		} else {
			t.Logf("Test case %d berhasil", i)
		}
	}
}
