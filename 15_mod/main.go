package main

import (
	"fmt"
	"itying/calc"
	T "itying/tools"
	"itying/tools/config"
)

func main() {
	sum := calc.Add(1, 2)
	fmt.Println(sum)
	fmt.Println(calc.Aaa)
	// fmt.Println(calc.aaa)
	fn1()

}

func fn1() {
	// tools.PrintInfo()
	T.PrintInfo()

	config.ConfigDB()
}

// init func run before main
func init() {
	fmt.Println("inited in main package...")
}
