package day3

import "testing"

func TestPrioForItem(t *testing.T) {

	type testInput struct {
		item    string
		expPrio int
	}

	tests := []testInput{
		{"a", 1},
		{"f", 6},
		{"z", 26},
		{"A", 27},
		{"Y", 51},
		{"Z", 52},
	}

	for _, tc := range tests {
		got := prioForItem(tc.item)
		if got != tc.expPrio {
			t.Errorf("Wrong priority for item %q, expected %d, got %d", tc.item, tc.expPrio, got)
		}
	}

}
