package main

import "fmt"

var x = 42

func main(){
	fmt.Println(x)
	test()
}

func test(){
	fmt.Println(x)
}