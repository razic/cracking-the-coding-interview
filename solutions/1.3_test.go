package main

import "testing"

func TestUrlifyEmptyString(t *testing.T) {
	strA, strB := []byte(""), []byte("")
	Urlify(strA, 0)
	if string(strA) != string(strB) {
		t.Fatalf("expected %q to be %q", strA, strB)
	}
}

func TestUrlifyASCIIString(t *testing.T) {
	strA, strB := []byte("Mr John Smith    "), []byte("Mr%20John%20Smith")
	Urlify(strA, 13)
	if string(strA) != string(strB) {
		t.Fatalf("expected %q to be %q", strA, strB)
	}
}
