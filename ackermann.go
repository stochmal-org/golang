package main

import (
	"fmt"
	"os"
	"strconv"
)

func acker(m int, n int) int {

	if m == 0 {
		return n + 1
	}

	if m > 0 && n == 0 {
		return acker(m-1, 1)
	}

	if m > 0 && n > 0 {
		return acker(m-1, acker(m, n-1))
	}

	return n + 1
}

func main() {

	m, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// ... handle error
		panic(err)
	}

	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		// ... handle error
		panic(err)
	}

	fmt.Printf("m = %d n = %d", m, n)

	var r = acker(m, n)

	fmt.Printf("\n%d", r)

}
