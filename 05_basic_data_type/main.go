package main

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/shopspring/decimal"
)

func main() {
	/**
	int
		int8, int16, int32, int64
		uint8, uint16, uint32, uint64

	float
		float32, float64

	bool: bool in go can't be operated on or converted to other types.
		true, false

	string
		"str"

	char
		unit8(int32), rune(UTF-8)

	rune
	*/

	var count int16 = 96
	fmt.Printf("count=%v, type: %T\n", count, count)
	fmt.Println(unsafe.Sizeof(count))

	var b float64 = 3.1415
	fmt.Printf("%v, %.2f, %T\n", b, b, b)
	fmt.Println(unsafe.Sizeof(b))

	d := 1129.6
	fmt.Println((d * 100))
	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2)
	d1 := decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(float64(m2)))
	fmt.Println(d1)

	e := float64(count)
	fmt.Printf("e=%v, types: %T\n", e, e)

	flag := true
	fmt.Printf("%v, %T\n", flag, flag)
	// a = 1
	// if a {
	// 	fmt.Println("xxx")
	// }

	str1 := "this is \na string variable"
	filePath := "//home//Larry//Projects"
	str3 := `1111
	22222
	3333`
	str4 := "11-22-33"

	fmt.Println(str1, len(str1))
	fmt.Println(filePath)
	fmt.Println(str3)
	fmt.Println(strings.Split(str4, "-"))

	var bt = 'c'
	var cc = '中'
	var bts = "this"
	var ct = "你想吃点啥"

	fmt.Printf("bt=%v, bt=%c, type: %T\n", bt, bt, bt)
	fmt.Printf("bt=%v, bt=%c, type: %T\n", cc, cc, cc)
	fmt.Printf("bts=%v, bts=%c, type=%T\n", bts[2], bts[2], bts[2])
	fmt.Printf("bts=%v, bts=%c, type=%T\n", ct[2], ct[2], ct[2])

	s := "你好golang"
	for _, v := range s { // rune
		fmt.Printf("%v(%c)\n", v, v)
	}

	strUpdate1 := "aabbddsd"
	byteStr := []byte(strUpdate1)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr))

	strUpdate2 := "皮条小猫"
	runeStr := []rune(strUpdate2)
	runeStr[3] = '狗'
	fmt.Println(string(runeStr))

}
