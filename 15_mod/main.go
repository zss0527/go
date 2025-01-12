package main

import (
	"fmt"
	"itying/calc"
	T "itying/tools"
	"itying/tools/config"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

func main() {
	sum := calc.Add(1, 2)
	fmt.Println(sum)
	fmt.Println(calc.Aaa)
	// fmt.Println(calc.aaa)
	fn1()
	thirdPartyPackageFn()
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

/*
1. import third party packages
2. go mod tidy to download related dependencies
*/
func thirdPartyPackageFn() {
	n, _ := decimal.NewFromString("-123.4567")
	fmt.Println(n.String()) // output: "-123.4567"

	json := `{
		"user":{
		   "first.name": "Janet",
		   "last.name": "Prichard"
		 }
	}`
	user := gjson.Get(json, "user")
	fmt.Println(user.Get(gjson.Escape("first.name")))
	fmt.Println(user.Get(gjson.Escape("last.name")))
	// Output:
	// Janet
	// Prichard
}
