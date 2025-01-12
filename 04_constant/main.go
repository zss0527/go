package main

import "fmt"

func main() {
	/*
		constant declaration and initialization can't be divided
		const canstant = expression
	*/

	const PI = 3.14
	const (
		A = "A"
		B = "B"
		C
		D
	)
	// const + iota +_
	const (
		AA = iota
		BB
		_
		CC
		DD
		E = iota
		F
	)
	fmt.Printf("PI = %v, type is %T\n", PI, PI)
	fmt.Println(A, B, C, D)
	fmt.Println(AA, BB, CC, DD, E, F)
}
