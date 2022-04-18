package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}

	fmt.Println(r.area())  // 50
	fmt.Println(r.perim()) // 30

	rp := &r
	fmt.Println(rp.area())  // 50
	fmt.Println(rp.perim()) // 30
}
