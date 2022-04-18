package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 24
	return &p
}

func main() {
	fmt.Println(person{"Ali", 18}) // {Ali 18}

	fmt.Println(person{name: "Veli", age: 20}) // {Veli 20}

	fmt.Println(person{name: "Ahmet"}) // {Ahmet 0}

	fmt.Println(&person{name: "Mehmet", age: 22}) // &{Mehmet 22}

	fmt.Println(newPerson("Hasan")) // &{Hasan 24}

	p := person{name: "Hüseyin", age: 26}
	fmt.Println(p.name) // Hüseyin

	pp := &p
	fmt.Println(pp.age) // 26

	pp.age = 27
	fmt.Println(pp.age) // 27
	fmt.Println(p.age)  // 27
}
