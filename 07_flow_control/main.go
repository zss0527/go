package main

import (
	"fmt"
)

func main() {
	/**
	if expression {
		code block
	} else if expression {
	 	code block
	} else {
	 	code block
	}

	for init; expression; end condition{
		code block
	}
	for end condition {
		code block
	}
	for {
		code block
	}
	for k, v := rang variable {
		code block
	}

	switch expression {
		case v1:
			code block
			break
		case v2:
			code block
			fallthrough
		case v3:
			code block
			break
		default:
			code block
	}
	switch {
		case c1:
			code block
			break
		case c2:
			code block
			break
		default:
			code block
	}

	no while loop in go

	continue, break, goto
	*/

	a := true

	if a {
		fmt.Println(a)
	}
	fmt.Println(a)

	/*
		variables declared in if expression, the variable scope is in the if.
	*/
	if b := 28; b > 18 {
		fmt.Println("ssssssss")
	}
	// fmt.Println(b)
	if score := 90; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	j := 1
	for j <= 3 {
		fmt.Println(j)
		j++
	}

	k := 1
	for {
		if k <= 5 {
			fmt.Println(k)
		} else {
			break
		}
		k++
	}

	str := "你好golang"
	for k, v := range str {
		fmt.Printf("key=%v, value=%c \n", k, v)
	}

	arr := []string{"php", "java", "golang", "ts"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	for k, v := range arr {
		fmt.Printf("key=%v, value=%v \n", k, v)
	}

}
