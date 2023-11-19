package main

import "fmt"

var a = "Check A"
var b, c string = " In B", " in C"
var d string

func main() {
	d = " in d"
	var e = 42
	f := 43
	g := " in g"
	h, i := "in h", "in i"
	j, k, l, m := 45.2, true, false, 'm'
	n := "n"
	o := `o`

	fmt.Println("a -", a)
	fmt.Println("b -", b)
	fmt.Println("c -", c)
	fmt.Println("d -", d)
	fmt.Println("e -", e)
	fmt.Println("f -", f)
	fmt.Println("g -", g)
	fmt.Println("h -", h)
	fmt.Println("i -", i)
	fmt.Println("j -", j)
	fmt.Println("k -", k)
	fmt.Println("l -", l)
	fmt.Println("m -", m)
	fmt.Println("n -", n)
	fmt.Println("o -", o)
}
