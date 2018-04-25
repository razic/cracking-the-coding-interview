package main

import "testing"

func TestIsUniqueEmptyString(t *testing.T) {
	str := ""

	if !IsUnique(str) {
		t.Fatalf("%x is expected to be unique", str)
	}
}

func TestIsUniqueASCIISingleChar(t *testing.T) {
	str := "a"

	if !IsUnique(str) {
		t.Fatalf("%q is expected to be unique", str)
	}
}

func TestIsUniqueASCIIWithDuplicates(t *testing.T) {
	str := "aaa"

	if IsUnique(str) {
		t.Fatalf("%q is expected to not be unique", str)
	}
}

func TestIsUniqueASCIIWithoutDuplicates(t *testing.T) {
	str := "abc"

	if !IsUnique(str) {
		t.Fatalf("%q is expected to be unique", str)
	}
}

func TestIsUniqueUnicodeSingleChar(t *testing.T) {
	str := "こ"

	if !IsUnique(str) {
		t.Fatalf("%q is expected to be unique", str)
	}
}

func TestIsUniqueUnicodeWithDuplicates(t *testing.T) {
	str := "こここ"

	if IsUnique(str) {
		t.Fatalf("%q is expected to not be unique", str)
	}
}

func TestIsUniqueUnicodeWithoutDuplicates(t *testing.T) {
	str := "こんに"

	if !IsUnique(str) {
		t.Fatalf("%q is expected to be unique", str)
	}
}
