package main

import (
	"fmt"
	"math"
)

const pi float32 = 3.14

func main() {
	fmt.Println(pi)

	const n = 500_000_000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}