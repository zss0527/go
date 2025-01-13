package main

import (
	"fmt"
	"reflect"
)

/*
reflect.TypeOf
reflect.ValueOf
*/
type myInt int
type myStruct struct {
	Name string
	Age  int
}

func main() {
	var e myInt = 13
	var f = myStruct{
		Name: "zhangsan",
		Age:  10,
	}
	kindFn(56)
	kindFn(24.5)
	kindFn(kindFn)
	kindFn(e)
	kindFn(f)
	kindFn(&e)
	kindFn([3]string{"aaa", "dddd"})
	kindFn([]string{"aaa", "dddd"})

	valueFn(56)
	valueFn(29.5)
}

func kindFn(x interface{}) {
	//get x type
	//one way is via x.(type)
	//another way is via reflect.TypeOf
	v := reflect.TypeOf(x)
	//Name() and Kind() methods in v
	fmt.Println(v, v.Name(), v.Kind())
}

func valueFn(x interface{}) {
	// b, _ := x.(int)
	// var num = b + 10
	// fmt.Println(num)

	v := reflect.ValueOf(x)
	//kind() in v
	t := v.Kind()
	switch t {
	case reflect.Int:
		fmt.Println(v.Int() + 10)
	case reflect.Float32:
		fmt.Println("float32")
	case reflect.String:
		fmt.Println("string type")
	default:
		fmt.Println("not find this type.")
	}

	// v.Elem().Kind()
	// v.Elem().SetInt(32)

}
