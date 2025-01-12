package main

import "fmt"

func main() {
	// var a = "variable"
	// fmt.Print("Hello Print!", "B", "c")
	// fmt.Println("hello Println", "A", "B", "C")
	// fmt.Printf("Hello Printf")
	// fmt.Printf("%v", a)

	/*
		the normal ways of variables definition in Golang
			1. var variable type = expression
				var name string = "Larry"
			2. var variable = expression, actual type is reffered by context
				var age = 28
			3.
	*/

	// var a int = 10
	// var b int = 5
	// var c int = 23
	a := 10
	b := 3
	c := 23

	fmt.Println("a=", a, "b=", b, "c=", c)
	fmt.Printf("a=%v b=%v c=%v\n", a, b, c)

}
