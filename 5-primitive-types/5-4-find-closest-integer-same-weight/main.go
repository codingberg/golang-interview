package main

import "fmt"

//seeking for any two consecutive bits that differ
func find_closest_integer(N uint64) uint64 {

	var pos uint
	for n := N; n > 0; n >>= 1 {
		if n&1 != n>>1&1 {
			return N ^ ((1 << pos) | (1 << (pos + 1)))
		}
		pos++
	}

	return N
}

//seeking for consecutive bits that will be transformed as 10->01
func find_closest_lower_integer(N uint64) uint64 {

	var pos uint
	for n := N; n > 0; n >>= 1 {
		if n&1 < n>>1&1 {
			return N ^ ((1 << pos) | (1 << (pos + 1)))
		}
		pos++
	}

	return N
}

//seeking for consecutive bits that will be transformed as 01->10
func find_closest_greater_integer(N uint64) uint64 {

	var pos uint
	for n := N; n > 0; n >>= 1 {
		if n&1 > n>>1&1 {
			return N ^ ((1 << pos) | (1 << (pos + 1)))
		}
		pos++
	}

	return N
}

func main() {

	fmt.Printf("Closest integer to %d: %d\n", 123, find_closest_integer(123))
	fmt.Printf("Closest lower integer to %d: %d\n", 123, find_closest_lower_integer(123))
	fmt.Printf("Closest greater integer to %d: %d\n", 123, find_closest_greater_integer(123))
}
