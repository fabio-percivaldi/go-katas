package main

import "testing"

func TestVowelCount(t *testing.T) {
	count := GetCount("abracadabra")

	if count != 5 {
		t.Fatalf("Expected 5, got %d", count)
	}
}
