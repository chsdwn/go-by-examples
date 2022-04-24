package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("42 cannot be processed")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "it cannot be processed"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{8, 42} {
		if r, err := f1(i); err != nil {
			fmt.Println("f1 failed:", err) // f1 worked: 11
		} else {
			fmt.Println("f1 worked:", r) // f1 worked: 42 - it cannot be processed
		}
	}

	for _, i := range []int{8, 42} {
		if r, err := f2(i); err != nil {
			fmt.Println("f2 failed:", err) // f2 worked: 11
		} else {
			fmt.Println("f2 worked:", r) // f2 failed: 42 - it cannot be processed
		}
	}

	_, err := f2(42)
	fmt.Println(err) // 42 - it cannot be processed
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)  // 42
		fmt.Println(ae.prob) // it cannot be processed
	}
}
