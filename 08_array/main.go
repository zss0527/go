package main

import "fmt"

func main() {
	/*
		Array declare
			var variable [length]type
				Array's length can't be change directly after declaration


		Array init
			1. arr[index] = value
			2. var variable = [len]type{value1,value2}
			   var variable = [...]type{value,value2}
			3. variable := [len]type{value1,value2}
			   variable := [...]type{value1,value2}
			4. variable := [...]type{index1:value1,index2: value2}


		Array loop
		1. for i := 0; i < len(arr); i++ {
			code block
		}
		2. for k, v := range arr {
			code block
		}

		multiple dimensional array
			var variable [len][len]type
			variable := [len][len]{{},{}}
			variable := [...][len]{{},{}}

		multiple dimensinal array loop
		for i := 0; i < len(arr[i]); i++ {
			for k := 0; k < len(arr[i][k]); k++ {
				code block
			}
		}
		for K1, v1 := range arr1 {
			for k2, v2 := range v2 {
				code block
			}
		}

	*/

	var arr1 [3]int
	var arr2 [4]string

	//arr1: [3]int, arr2: [4]string
	//so, [len]T is one kind of type in go
	fmt.Printf("arr1: %v, type: %T\narr2: %v, type: %T\n", arr1, arr1, arr2, arr2)

	arr1[2] = 4
	fmt.Println(arr1)

	arr3 := [...]string{"c", "java", "go", "js/ts", "python"}
	fmt.Println(arr3)
	arr4 := [...]int{1: 33, 4: 23, 19: 22}
	fmt.Println(len(arr4))

	var arr5 = [3][2]string{
		{"beiing", "shanghai"},
		{"guzhou", "shenzhen"},
		{"chengdu", "chognqing"},
	}
	fmt.Println(arr5)

	//basic type and Array are value type
	a := 1
	b := a
	a = 2
	fmt.Println(a, b)

	ar1 := [...]int{1, 2, 3}
	ar2 := ar1
	ar1[2] = 7
	fmt.Println(ar1, ar2)

	//slice is reference type
	sl1 := []int{1, 2, 3}
	sl2 := sl1
	sl1[2] = 7
	fmt.Println(sl1, sl2)

}
