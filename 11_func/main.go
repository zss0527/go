package main

import (
	"errors"
	"fmt"
	"sort"
)

/*
1. func declare
	func funcName(params)(return types) {
		code block
	}

2. variable scope
	block scope variable: declared in scope(func,if,for. etc), can't used out of this scope.
	global scope variable: declared out of main, can be used anywhere.

3. func type, func as params, fun as return, anonymous func
	type yourDefinedName func(params types)(return types)
		type calc func(int,int) int  // self defined calc type

	the value of func name is the starting address of the func body,
	unlike other variable's type(int8 is a scope), type of func name is just the func signature.
	when assigning function names, their types must be the same
	the essence of func call(funcName()) is run func body code and copy the actual params to the formal params in func body

	func as param
		func funcAsParam(x, y int, fn func(a, b string) string) {}
		func funcAsParam(x, y int, fn selfDefinedFuncType) {}

	fun as return
		func funcAsReturn(x, y int) func(a, b string) string {}
		func funcAsReturn(x, y int) selfDefinedFuncType {}

	anonymous function
		anoFn := func(x, y int) int {}

4. recursive call, closure function
	func outer() func() int {
		var i = 10
		return func() int {
			return i + 1
		}
	}

5. defer: will be executed between return value assignment and return directive
	return value assignment
	run defer
	return directive

6. panic/recover package


*/

func main() {
	test()
	sum := sumFn(12, 13)
	fmt.Println(sum)
	fmt.Println(subFn(4, 7))
	fmt.Println(multiplParamsFn(1, 2, 3, 4))

	a, b := multipleReturnFn(10, 3)
	fmt.Println(a, b)
	fmt.Println(multipleReturnFn(10, 3))

	c, _ := multipleReturnNameFn(21, 3)
	fmt.Println(c)

	fmt.Println(multipleReturnNameFn(21, 3))

	arr := []int{3, 2, 7, 7, 4, 1, 9}
	sortIntAsc(arr)
	fmt.Println(arr)

	m1 := map[string]string{
		"username": "Larry",
		"age":      "28",
		"sex":      "male",
		"height":   "175",
	}
	fmt.Println(mapSort(m1))

	varScopeFun()

	var cc, dd calc
	cc = sumFn
	// type of test is not calc
	//cc = test
	dd = sumFn
	fmt.Printf("cc type: %T,  cc: %v\n", cc, cc)
	fmt.Printf("dd type: %T,  dd: %v\n", dd, dd)

	fmt.Printf("sumFn type: %T, sumFn: %v\n", sumFn, sumFn)
	fmt.Println(cc(4, 5))

	var myA myInt = 10
	var mya int = 20
	//can't add directly
	// var myb = myA + mya
	fmt.Printf("myA type: %T, myA: %v\n", myA, myA)
	fmt.Printf("mya type: %T, mya: %v\n", mya, mya)

	funcAsParam1(1, 2, sumFn)
	funcAsParam2(1, 2, sumFn)
	//anonymous function
	funcAsParam2(5, 7, func(x, y int) int {
		return x * y
	})
	anoFn := func(x, y int) int {
		return x / y
	}
	anoFn(8, 4)
	//self execution of anomymous function
	func(x, y int) {
		fmt.Println("anonymous function in main.", x, y, x*y)
	}(7, 8)

	fmt.Println(recrusiveFn(100))
	fmt.Println(factorialFn(5))

	//closure func
	var fn1 = outer1()
	fmt.Println(fn1())
	fmt.Println(fn1())
	fmt.Println(fn1())

	var fn2 = outer2()
	fmt.Println(fn2(3))
	fmt.Println(fn2(3))
	fmt.Println(fn2(3))

	//defer
	// fmt.Println("start...")
	// fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("finished...")
	deferFn1()
	fmt.Println(deferFn2())
	fmt.Println(deferFn3())
	fmt.Println(deferFn4())
	fmt.Println(deferFn5())
	fmt.Println(deferFn6())

	//panic
	paFn1()
	paFn2()
	fmt.Println("test panic ...")
	paFn3(10, 0)
	fmt.Println(paFn3(10, 2))
	myFn()
}

func test() {
	fmt.Println("no params, no return")
}

func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

// simplify func params
func subFn(x, y int) int {
	return x - y
}

func multiplParamsFn(y int, x ...int) int {
	fmt.Printf("y:%v,------x:%v,------x type:%T\n", y, x, x)

	var sum int
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func multipleReturnFn(x, y int) (int, int) {
	sum := x + y
	sub := x - y

	return sum, sub
}

// or func multipleReturnNameFn(x, y int) (sum, sub int)
func multipleReturnNameFn(x, y int) (sum int, sub int) {
	fmt.Println(sum, sub)
	sum = x * x
	sub = x / y

	return
}

func sortIntAsc(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
}

func mapSort(map1 map[string]string) string {
	var sliceKey []string

	for k, _ := range map1 {
		sliceKey = append(sliceKey, k)
	}

	sort.Strings(sliceKey)

	var str string
	for _, v := range sliceKey {
		str += fmt.Sprintf("%v=>%v", v, map1[v])
	}
	return str
}

var glVar = "global variable"

// global variable can't be declared via :=
// glVarByColon := "dddddd"
func varScopeFun() {
	//block scope variable
	var glVar = "11111"
	fmt.Println(glVar)
}

type calc func(int, int) int
type myInt int

func funcAsParam1(x, y int, c func(int, int) int) {
	fmt.Println("c(x,y) in funcAsParam1: ", c(x, y))
}

func funcAsParam2(x, y int, c calc) {
	fmt.Println("c(x,y) in funcAsParam2: ", c(x, y))
}

func funcAsReturn(operator string) calc {
	switch operator {
	case "+":
		return sumFn
	case "-":
		return subFn
	default:
		return nil
	}
}

// function recursive call
func recrusiveFn(n int) int {
	if n > 1 {
		// fmt.Println(n)
		// n--
		// recrusiveFn(n)
		return n + recrusiveFn(n-1)
	} else {
		return 1
	}
}

func factorialFn(n int) int {
	if n > 1 {
		return n * factorialFn(n-1)
	} else {
		return 1
	}
}

// closure func
func outer1() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func outer2() func(x int) int {
	//i is in memory always like global scope variable, but i have no global influence.
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

// defer selft execution func
func deferFn1() {
	fmt.Println("start...")
	defer func() {
		fmt.Println("aaa")
	}()
	fmt.Println("finished...")
}

func deferFn2() int { //0
	var a int //0
	defer func() {
		a++
		fmt.Println("a in deferFn2", a)
	}()
	return a
}

func deferFn3() (a int) { //11
	defer func() {
		a++
		fmt.Println("a in deferFn3", a)
	}()
	return 10
}

func deferFn4() (x int) { //5
	y := 5
	defer func() {
		y++
	}()
	return y
}

func deferFn5() (y int) { //6
	y = 5
	defer func() {
		y++
	}()
	return y
}

// defer func with params case
func deferFn6() (x int) { //
	defer func(x int) {
		fmt.Println("x in deferFn6", x)
		x++
		fmt.Println("x in deferFn6", x)
	}(x) //value of defer func params must be confirmed when defer registered the func, so the x value is 0 when the func registered
	return 5
}

// panic/recover
func paFn1() {
	fmt.Println("paFn1")
}
func paFn2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("find error in paFn2 and catched")
		}
	}()
	panic("throw one exception")
}

func paFn3(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	return a / b
}

func readMain(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("read file failed!")
	}
}
func myFn() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("erroes occured while read file:", e)
		}
	}()

	err := readMain("demo.go")
	if err != nil {
		panic(err)
	}
}
