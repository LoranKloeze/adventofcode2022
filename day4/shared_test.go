package day4

import (
	"testing"
)

func TestExtractPairs(t *testing.T) {

	tests := []struct {
		input    string
		expPair1 pair
		expPair2 pair
	}{
		{"2-4,6-8", pair{2, 4}, pair{6, 8}},
		{"3-3,6-46", pair{3, 3}, pair{6, 46}},
		{"7-8,56-897", pair{7, 8}, pair{56, 897}},
	}

	for _, tc := range tests {

		gotPair1, gotPair2, err := extractPairs(tc.input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if gotPair1 != tc.expPair1 {
			t.Errorf("Expected %v to be extracted as pair 1 from %q, got %v ", tc.expPair1, tc.input, gotPair1)
		}
		if gotPair2 != tc.expPair2 {
			t.Errorf("Expected %v to be extracted as pair 2 from %q, got %v ", tc.expPair2, tc.input, gotPair2)
		}

	}
}

func TestExtractPairsWithNonNumbers(t *testing.T) {

	tests := []string{
		"abc-4,5-8",
		"5-wha,8-12",
		"cc-wha,e-1r",
	}

	for _, tc := range tests {

		_, _, err := extractPairs(tc)

		if err == nil { // err IS nil
			t.Fatalf("Expected error extracting pairs from %q, got none", tc)
		}

	}
}
