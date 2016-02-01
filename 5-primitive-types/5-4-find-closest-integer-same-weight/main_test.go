package main

import (
	"testing"
)

var closestIntegerTests = []struct {
	n              uint64
	expected       uint64
	lower, greater uint64
}{
	{123, 125, 119, 125},
	{92, 90, 90, 108},
	{6, 5, 5, 10},
	{2, 1, 1, 4},
	{32, 16, 16, 64},
	{0xFFFFFFFE, 0xFFFFFFFD, 0xFFFFFFFD, 0x17FFFFFFE},
}

func TestSet(t *testing.T) {
	for _, tt := range closestIntegerTests {
		expected := find_closest_integer(tt.n)
		if expected != tt.expected {
			t.Errorf("find_closest_integer(%d) actual %d != expected %d", tt.n, expected, tt.expected)
		}
		expected = find_closest_lower_integer(tt.n)
		if expected != tt.lower {
			t.Errorf("find_closest_lower_integer(%d) actual %d != expected %d", tt.n, expected, tt.lower)
		}
		expected = find_closest_greater_integer(tt.n)
		if expected != tt.greater {
			t.Errorf("find_closest_greater_integer(%d) actual %d != expected %d", tt.n, expected, tt.greater)
		}
	}
}

func BenchmarkFindGreaterThan_FFFE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		find_closest_greater_integer(0xFFFE)
	}
}
func BenchmarkFindGreaterThan_FFFFFE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		find_closest_greater_integer(0xFFFFFE)
	}
}

func BenchmarkFindGreaterThan_FFFFFFFE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		find_closest_greater_integer(0xFFFFFFFE)
	}
}

func BenchmarkFindLowerThan_FFFE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		find_closest_lower_integer(0xFFFE)
	}
}

func BenchmarkFindLowerThan_FFFFFE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		find_closest_lower_integer(0xFFFFFE)
	}
}

func BenchmarkFindLowerThan_FFFFFFFE(b *testing.B) {
	for n := 0; n < b.N; n++ {
		find_closest_lower_integer(0xFFFFFFFE)
	}
}
