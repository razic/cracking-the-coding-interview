package main

import "testing"

func TestCheckPermutationEmptyStrings(t *testing.T) {
	strA, strB := "", ""
	if CheckPermutation(strA, strB) {
		t.Fatalf("%q %q should not be a permutation", strA, strB)
	}

	strA, strB = "hello", ""
	if CheckPermutation(strA, strB) {
		t.Fatalf("%q %q should not be a permutation", strA, strB)
	}

	strA, strB = "", "hello"
	if CheckPermutation(strA, strB) {
		t.Fatalf("%q %q should not be a permutation", strA, strB)
	}
}

func TestCheckPermutationASCII(t *testing.T) {
	strA, strB := "hello", "hello"
	if CheckPermutation(strA, strB) {
		t.Fatalf("%q %q should not be a permutation", strA, strB)
	}

	strA, strB = "helol", "hello"
	if !CheckPermutation(strA, strB) {
		t.Fatalf("%q %q should not be a permutation", strA, strB)
	}
}

func TestCheckPermutationUnicode(t *testing.T) {
	strA, strB := "こんに", "こんに"
	if CheckPermutation(strA, strB) {
		t.Fatalf("%q %q should not be a permutation", strA, strB)
	}

	strA, strB = "こんに", "こにん"
	if !CheckPermutation(strA, strB) {
		t.Fatalf("%q %q should not be a permutation", strA, strB)
	}
}
