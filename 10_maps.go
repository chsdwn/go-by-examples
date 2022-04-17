package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 8
	m["k2"] = 24
	fmt.Println("map:", m)
	fmt.Println("length:", len(m))

	k1 := m["k1"]
	fmt.Println("k1:", k1)

	delete(m, "k2")
	fmt.Println("map:", m)

	_, ok := m["k2"]
	fmt.Println("k2 ok:", ok)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("initialized:", n)
}
