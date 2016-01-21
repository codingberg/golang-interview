package main

import (
	"math/rand"
	"testing"
)

type reverseTest struct {
	n        uint64
	expected uint64
}

var reverseTests []reverseTest

func init() {
	for i := 0; i < 8196; i++ {
		number := uint64(rand.Int63())
		reverseTests = append(reverseTests, reverseTest{number, divreverse(number)})
	}
}

func TestBruteReverse(t *testing.T) {
	for _, tt := range reverseTests {
		actual := brutereverse(tt.n)
		if actual != tt.expected {
			t.Errorf("BruteReverese(%x) expected: %x actual: %x ", tt.n, tt.expected, actual)
		}
	}
}

func TestLookupReverse(t *testing.T) {
	for _, tt := range reverseTests {
		actual := lookupreverse(tt.n)
		if actual != tt.expected {
			t.Errorf("LookupReverese(%x) expected: %x actual: %x ", tt.n, tt.expected, actual)
		}
	}
}

func BenchmarkDivReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divreverse(1234565789)
	}
}

func BenchmarkBruteReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		brutereverse(1234565789)
	}
}

func BenchmarkLookupReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lookupreverse(1234565789)
	}
}
