package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	/*
		函数用来抽离代码、复用、切分大的功能为小的功能、向调用者隐藏实现细节
	*/
	f := square
	fmt.Println(f(3))
	// f = product  //不同类型的函数变量不能相互赋值

	//函数类型的零值为nil（空值），调用一个空函数会导致宕机
	var t func(int) int
	// t(3)//会导致宕机

	//函数变量可以和空值相比较
	if t != nil {
		t(3)
	}

	//函数变量使得函数不仅将数据参数化，还将函数的行为当作参数进行传递
	//strings.Map对字符串中的每一个字符使用一个函数，将结果连接起来变成另一个字符串
	fmt.Println(strings.Map(add1, "VMS")) //"WNT"
	//匿名函数能够获取到整个词法环境，内层的函数可以访问外层函数里的变量
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "IAM"))

	//ff能够获取和更新外层squares函数的局部变量，由于这些隐藏的变量引用所以把函数定义为引用类型而且函数变量没法进行比较
	//函数变量类似于闭包
	//函数局部变量的生命周期不是由它的作用域决定的：变量x在main函数中返回squares函数后依旧存在（被ff引用）
	ff := squares()
	fmt.Println(ff()) //"1"
	fmt.Println(ff()) //"4"

	sum(1, 2, 3, 4, 5)
	sum([]int{1, 2, 3}...)

}

//函数声明和声明其它类型变量一样，每个函数声明包含一个名字、一个形参列表、一个可选的返回列表以及函数体
// func name(parameter-list) (result-list) {Body}
//函数也和其他变量一样拥有类型叫做函数签名，函数签名包括函数名、形参列表和返回列表，其中参数和返回列表的顺序是签名的一部份
//返回值也可以像形参一样命名（通常为指针类型），这时候每一个命名的返回值会声明为一个局部变量，函数return时可以不包含任何内容
//形参和返回列表（如果是命名的话）是该函数的局部变量，初始值为类型零值，被调用时由实参赋值
//实惨是按值传递的，如果提供的实参包含引用类型（指针、slice、map、函数、channel）那么函数修改对应的形参会间接影响外面原来的值

//一个函数能够返回多个结果通过多个变量来接收，如果某些返回值不会被用到可以用空标识符"_"接收

// 函数变量也有类型，而且他们可以赋给便利那个或者传递或者从其他函数中返回
func square(n int) int     { return n * n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }

//命名函数只能在包级别的作用域进行声明，但我们能够使用函数字面量在任何表达式内指定函数变量
//函数字面量就像函数声明，但在func关键字后面没有函数的名称，它是一个表达式，它的值称做匿名函数

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// 变长函数被调用的时候可以有可变的参数个数
// 在参数列表最后的类型名称之前使用省略号"..."表示声明一个变长函数，调用这个函数的时候可以传递该类型任意数目的参数
func sum(vals ...int) int { //在函数体内部变长参数是个int类型的slice
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//sum(1,2,3,4)
//如果实参已经是slice类型，需要在传递的时候在后面加省略号"..."
//sum([]int{1,2,3,4}...)

// 延迟函数调用是指在函数调用之前加上关键字defer，会在函数return之前调用，defer使用次数没有要求，多个defer执行顺序是倒叙进行
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// go语言中的宕机类似于java中的运行时异常，除了系统在运行时遇到错误触发宕机以外，也可以在程序中主动调用内置的宕机函数panic
func Compile(expr string) (string, error) { return "", errors.New("something wrong") }
func MustCompile(expr string) string {
	re, err := Compile(expr)
	if err != nil {
		panic(err)
	}
	return re
}

//出现宕机时除了退出程序外有时候也可以进行恢复，内置的recover函数结合defer可以恢复宕机，类似于java中的try-catch
