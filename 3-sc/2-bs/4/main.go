package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increament := wrapper()
	fmt.Println(increament())
	fmt.Println(increament())
}
