package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println("1:", nextInt())
	fmt.Println("1:", nextInt())
	fmt.Println("1:", nextInt())

	nextInt2 := intSeq()
	fmt.Println("2:", nextInt2())
}
