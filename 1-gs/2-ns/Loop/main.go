package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d %b %#x \n", i, i, i)
	}
}
