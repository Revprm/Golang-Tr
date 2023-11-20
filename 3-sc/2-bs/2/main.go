package main

import "fmt"

var x = 0

func increament() int{
	x++
	return x
}

func main(){
	fmt.Println(increament())
	fmt.Println(increament())
}