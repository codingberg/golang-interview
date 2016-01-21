package main

import (
	"fmt"
	"log"
	"strconv"
)

// swap bits if they differ
func swapbits(n uint64, i uint, j uint) uint64 {
	if n>>i&1 != n>>j&1 {
		n ^= (1 << i) | (1 << j)
	}
	return n
}

func main() {
	var n uint64
	var i, j uint

	fmt.Printf("Enter number and bit indicies [n i j]: ")
	_, err := fmt.Scanf("%d %d %d", &n, &i, &j)
	if err != nil {
		log.Fatal(err)
	}

	bn := strconv.FormatUint(n, 2)
	sn := swapbits(n, i, j)
	bsn := strconv.FormatUint(sn, 2)

	fmt.Printf("%8s\n%8s\n", bn, bsn)
}
