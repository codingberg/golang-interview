package main

import "testing"

//some simple tests
var swapTests = []struct {
	n        uint64
	i, j     uint
	expected uint64
}{
	// @boundaries
	{0x1, 7, 0, 0x80},
	{0x1, 15, 0, 0x8000},
	{0x1, 31, 0, 0x80000000},
	{0x1, 47, 0, 0x800000000000},
	{0x1, 63, 0, 0x8000000000000000},
	// all bits are the same
	{0xff, 5, 2, 0xff},
	// ordinary different sizes integers + @boundaries
	{0xff, 15, 2, 0x80fb},
	{0xfb, 5, 2, 0xdf},
	{0xfffb, 13, 2, 0xdfff},
	{0xfffffffb, 29, 2, 0xdfffffff},
	{0xfffffffffffffffb, 61, 2, 0xdfffffffffffffff},
	{0xfe, 7, 0, 0x7f},
	{0xfffe, 15, 0, 0x7fff},
	{0xfffffffe, 31, 0, 0x7fffffff},
	{0xfffffffffffffffe, 63, 0, 0x7fffffffffffffff},
}

func TestAll(t *testing.T) {
	for _, tt := range swapTests {
		actual := swapbits(tt.n, tt.i, tt.j)
		if actual != tt.expected {
			t.Errorf("swapbits(%x,%d,%d): expected %x, actual %x", tt.n, tt.i, tt.j, tt.expected, actual)
		}

	}
}
