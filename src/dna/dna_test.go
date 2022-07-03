package dna

import "testing"

func TestDna(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"AAAA", "TTTT"},
		{"ATTGC", "TAACG"},
		{"GTAT", "CATA"},
	}

	for _, tc := range testCases {
		if got := DNAStrand(tc.input); got != tc.expected {
			t.Errorf("Expected %s and got %s", tc.expected, got)
		}
	}
}
