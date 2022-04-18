package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "Ali",
	}

	fmt.Println(co.num, co.str) // 1 Ali

	fmt.Println(co.base.num) // 1

	fmt.Println(co.describe()) // base with num=1

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println(d.describe()) // base with num=1
}
