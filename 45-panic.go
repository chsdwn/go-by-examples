package main

import "os"

func main() {
	panic("an error occured")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
