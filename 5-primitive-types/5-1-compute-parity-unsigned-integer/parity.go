package main

import (
	"fmt"
	"strconv"
	"strings"
)

func f1(n uint64) int {
	bs := strconv.FormatUint(n, 2)
	return strings.Count(bs, "1")
}

func f2(n uint64) int {
	cnt := 0
	for n > 0 {
		if n&1 == 1 {
			cnt++
		}
		n = n >> 1
	}
	return cnt
}

func f3(n uint64) int {
	cnt := 0
	for n > 0 {
		n &= (n - 1)
		cnt++
	}
	return cnt
}

var bitCounts = []int8{
	0, 1, 1, 2, 1, 2, 2, 3,
	1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4,
	2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4,
	2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5,
	3, 4, 4, 5, 4, 5, 5, 6,
	1, 2, 2, 3, 2, 3, 3, 4,
	2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5,
	3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5,
	3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6,
	4, 5, 5, 6, 5, 6, 6, 7,
	1, 2, 2, 3, 2, 3, 3, 4,
	2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5,
	3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5,
	3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6,
	4, 5, 5, 6, 5, 6, 6, 7,
	2, 3, 3, 4, 3, 4, 4, 5,
	3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6,
	4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6,
	4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7,
	5, 6, 6, 7, 6, 7, 7, 8,
}

func f4(n uint64) int {
	var cnt int8
	cnt += bitCounts[n&255]
	cnt += bitCounts[n>>8&255]
	cnt += bitCounts[n>>16&255]
	cnt += bitCounts[n>>24&255]
	cnt += bitCounts[n>>32&255]
	cnt += bitCounts[n>>40&255]
	cnt += bitCounts[n>>48&255]
	cnt += bitCounts[n>>56&255]
	return int(cnt)
} // end-of-f4

func main() {
	x := uint64(123456789)
	fmt.Println("Binary ", strconv.FormatUint(x, 2))

	fmt.Println("f1 =", f1(123456789))
	fmt.Println("f2 =", f2(123456789))
	fmt.Println("f3 =", f3(123456789))
	fmt.Println("f4 =", f4(123456789))
}
