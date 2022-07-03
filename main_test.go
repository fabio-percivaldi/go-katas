package main

import "testing"

func TestVowelCount(t *testing.T) {
	count := GetCount("abracadabra")

	if count != 5 {
		t.Fatalf("Expected 5, got %d", count)
	}

	count = GetCount("ala")

	if count != 2 {
		t.Fatalf("Expected 2, got %d", count)
	}

	count = GetCount("ape")

	if count != 2 {
		t.Fatalf("Expected 2, got %d", count)
	}
}
