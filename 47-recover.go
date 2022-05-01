package main

import "fmt"

func mayPanic() {
	panic("an error occured")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")

	/* OUTPUT */
	// Recovered. Error:
	//	an error occured
}
