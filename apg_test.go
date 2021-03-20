package main

import "testing"

// Make sure the flags are initalized
var _ = func() bool {
	testing.Init()
	return true
}()

// Test getRandNum with max 1000
func TestGetRandNumMax1000(t *testing.T) {
	randNum := getRandNum(1000)
	if randNum > 1000 {
		t.Errorf("Generated random number between 0 and 1000 is too big: %d", randNum)
	}
	if randNum < 0 {
		t.Errorf("Generated random number between 0 and 1000 is too small: %d", randNum)
	}
}

// Test getRandNum with max 1
func TestGetRandNumMax1(t *testing.T) {
	randNum := getRandNum(1)
	if randNum > 1 {
		t.Errorf("Generated random number between 0 and 1 is too big: %d", randNum)
	}
	if randNum < 0 {
		t.Errorf("Generated random number between 0 and 1 is too small: %d", randNum)
	}
}

// Test getRandChar
func TestGetRandChar(t *testing.T) {
	charRange := "ABC"
	randChar := getRandChar(&charRange, 1)
	if randChar != "A" && randChar != "B" && randChar != "C" {
		t.Errorf("Random character generation failed. Expected A, B or C but got: %v", randChar)
	}

	randChar = getRandChar(&charRange, 1000)
	if len(randChar) != 1000 {
		t.Errorf("Generated random characters with 1000 chars returned wrong amount of chars: %v", len(randChar))
	}
}

// Benchmark: Random number generation
func BenchmarkGetRandNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRandNum(100000)
	}
}

// Benchmark: Random char generation
func BenchmarkGetRandChar(b *testing.B) {
	charRange := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890\"#/!\\$%&+-*.,?=()[]{}:;~^|"
	for i := 0; i < b.N; i++ {
		getRandChar(&charRange, 20)
	}
}
