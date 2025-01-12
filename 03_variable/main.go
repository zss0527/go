package main

import "fmt"

var global_variable = "globale variable"

// function_variable := "function variable"
//variable := expression, this way just can be used in function, can't decalre as global

func getUserInfo() (string, int) {
	return "zhangsan", 10
}

func main() {
	/*
		variable declaration and initialization
		1. just declare
		var variable type
			var userName string
			var (
				userName string
				age int
				address string
			)

		2. declare and init
			2.1. var variable type = expression
					var userName string = "Larry"
			2.2. var variable = expression
					var userName = "Larry"
			2.3. variable := expression
					userName := "Larry"
	*/

	var (
		userName string
		age      int
		address  string
	)

	userName = "Larry"
	age = 28
	address = "China"

	fmt.Printf("userName=%v, age=%v, address=%v\n", userName, age, address)

	var (
		hobby string = "game"
		sex   string = "male"
	)

	fmt.Printf("hobby=%v, sex=%v\n", hobby, sex)

	color := "red"
	a, b, c := 10, "weeeee", 23

	fmt.Printf("color = %v\n", color)
	fmt.Printf("a=%v type is %T, b=%v type is %T, c=%v type is %T\n", a, a, b, b, c, c)

	/*
		anonymous variable
	*/
	var username, userage = getUserInfo()
	// just get username, second name ignore
	var justUserName, _ = getUserInfo()
	fmt.Printf("username=%v, userage=%v\n", username, userage)
	fmt.Printf("justUserName = %v", justUserName)

}
