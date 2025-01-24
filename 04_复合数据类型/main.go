package main

import "fmt"

/*
go中的复合数据类型是由基本数据类型以各种方式组合而构成的，主要包括数组、slice、map和struct
其中数组和结构体是聚合类型，它们的值由内存中的一组变量构成，数组的元素具有相同的类型，结构体的元素数据类型可以不同；
数组和结构体长度固定，slice和map是动态数据结构，长度可以动态增长；
*/

func main() {
	// arrayFn()
	sliceFn()

}

func arrayFn() {
	/*
		由于数组长度固定，所以在实际应用中使用更多的是长度可变的slice，但是数组是slice的基础（底层类型）
	*/
	//1.1 数组中的每个元素只能通过索引来访问，索引值从0开始，内置的len函数返回数组的元素个数；
	var a [3]int
	fmt.Println(a[0])        //first element
	fmt.Println(a[len(a)-1]) //last element
	//通常使用for range遍历数组
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	//仅输出元素
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	//默认情况下数组元素初始为类型零值，也可以通过数组字面量初始化一个数组
	var q [3]int = [3]int{3, 4, 6}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q, r[2]) //q,0
	//在数组字面量中如果省略号...出现在数组长度位置，那么数组长度由初始化数组的元素个数决定
	m := [...]int{1, 2, 3}
	fmt.Printf("%T\n", m)
	//数组的长度是数组类型的一部份，所以[3]int和[4]int是两种不同的数组类型
	//数组的长度必须是常量表达式，表达式的值在编译期就得确定
	t := [3]int{1, 2, 3}
	fmt.Println(t)
	//t = [4]int{1,2,3,4}//错误，不能将[4]int复制给[3]int
	//数组的一种用法
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "&", GBP: "@", RMB: "¥"}
	fmt.Println(symbol)
	//如果数组元素类型是可比较的，那么这个数组整体也是可比较的
	a1 := [2]int{1, 2}
	a2 := [...]int{1, 2}
	a3 := [2]int{1, 3}
	fmt.Println(a1 == a2, a1 == a3, a2 == a3) //true false false
	a4 := [3]int{1, 2}
	// fmt.Println(a1 == [2]int(a4)//编译错误，无法比较[2]int和[3]int
	fmt.Println(a4)
	//由于数组是值类型而非引用类型，所以接收数组的函数会创建新的副本，而不是直接操作原始的数组
	//传递大数组是低效的，可以采用传递数组指针的方式传递大数组，这样函数内部对数组的任何修改都会反映到原始数组上；
	/*
		func zero(ptr *[32]byte) {
			for i := range ptr {
				ptr[i] = i
			}
		}
	*/
	//使用数组指针是高效的，但因为数组本身长度固定所以常常使用slice替代
}

func sliceFn() {
	/*
		slice表示一个拥有相同类型元素的可变长度序列，slice通常写成[]T，元素类型是T；
		slice底层总是绑定一个数组，这个数组称为slice的底层数组，可以通过slice访问这个底层数组的部分或全部元素；
		slice有三个属性：指针、长度和容量
			指针指向数组的第一个可以从slice中访问的元素；
			长度是指slice中元素个数，它不能超过slice的容量；
			容量指从slice的起始元素到底层数组的最后一个元素间元素的个数；
			go内置的len和cap用来返回slice的长度和容量；
	*/
	//一个底层数组可以对应多个slice，这些slice可以引用数组的任何位置，彼此之间的元素还可以重叠
	months := [...]string{0: "ZeroMonth",
		1:  "January", //基于months声称接下来的一些列slice，months称为这些slice的底层数组
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December"}
	//slice操作赋s[i:j](0<=j<=j<=cap(s))创建了一个新的slice，这个新的slice引用了序列s中从i到j-1索引位置的所有元素
	//这里的s可以是数组，或者指向数组的指针，也可以是slice
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2, summer) //["April" "May" "June"],["June" "July" "August"]
	//Q2和summer共享了June，通过一个slice修改底层数组另一个也会受影响
	Q2[2] = "Test"
	fmt.Println(Q2, summer) ////["April" "May" "Test"],["Test" "July" "August"]

	//如果slice的引用超过了被引用对象的容量即cap(s)会导致宕机，如果导致超出了被引用对象的长度即len(s)会导致slice扩容
	// fmt.Println(summer[:20])   //宕机
	endlessSummer := summer[:5] //导致slice容量扩容
	fmt.Println(endlessSummer)  //["June" "July" "August" "September" "Octover"]

	//求字符串子串操作和对字节slice（[]byte）做slice操作原理相同，但字符串子操作返回子串，字节slice子操作返回slice

	//slice作为函数参数时传递的是引用，所以可以通过函数就地修改slice和它的底层数组
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	reverse(a[:])
	fmt.Println(a)

	/*
		和数组不同的是slice无法做比较，因此不能用==来测试两个slice是否拥有相同的元素
		由于散列表map仅对键做浅拷贝，要求散列表的键在散列表的整个生命周期内必须保持不变，这就要求要对键做深度比较
		所以不能用slice作为map的键
		对于引用类型pointer和channel，==比较的是引用相等性
		唯一可以与slice直接比较的是nil，slice的零值是nil，值为nil的slice没有对应的底层数组且len和cap均为0
		如果要检查一个slice是否为空，使用len(s) == 0，而不是s == nil，因为s != nil的情况下slice也有可能是空
	*/

	//内置的make可以创建一个具有指定元素类型、长度和容量的slice，其中容量参数可以省略此时表示该slice容量和长度相同
	ms := make([]int, 10)
	fmt.Println(len(ms), cap(ms))
	//其实make创建了一个无名数组并返回了它的一个slice，且这个数组仅可以通过这个slice访问

	//内置的append函数用来将元素追加到slice的后面
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r) //尽管slice是引用类型，但使用append函数为slice追加元素时仍然需要重新接收结果
	}
	fmt.Printf("%q\n", runes) //['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
	//appendInt展示了原因
	an := appendInt(a[:], 7)
	fmt.Println(an)
	//所以，对于任何函数，只要有可能改变slice的长度或者容量，抑或是使得slice指向不同的底层数组，都需要重新更新slice变量
	/*
		为了正确地使用slice，必须记住，虽然底层数组的元素是间接引用的，但是slice的指针、长度和容量不是
		要更新一个slice的指针、长度或容量必须使用类似appendInt中make显式赋值
		所以， slice并不是一个纯引用类型，而是像下面这种聚合类型
		type IntSlice struct {
			ptr *int  //指向底层数组第一个元素，所以同一个底层数组的多个slice的指针是相同的，即多个&slice[0]相同，但多个slice[0]即首个元素可能不一样
			len, cap int
		}
		slice的地址&slice获取的是这个结构体的地址
		&slice[0]获取到的是底层数组的第一个元素的地址
		slice[0]获取到的是底层数组中该slice可以访问到的第一个元素
	*/

	//slice可以用来实现栈，给定一个空的slice元素stack，可以使用append向slice尾部追加值
	var stack []int
	stack = append(stack, 1) //push 1
	stack = append(stack, 2) //push 2
	//栈的顶部是最后一个元素
	top := stack[len(stack)-1]
	fmt.Println(top)
	stack = stack[:len(stack)-1] //pop

	/*
		总结创建slice的方式
	*/
	//1. slice字面量直接初始化
	slice1 := []int{1, 2, 3}
	//2. 通过make函数创建一个指定长度和容量的slice
	slice2 := make([]int, 2, 3)
	//3. 对已有数组或者slice通过slice操作获取新的slice
	slice3 := a[3:5]
	//4. 声明一个slice以后通过append追加元素，这种刚声明的slice不能直接通过下标赋值，因为还没有分配空间
	var slice4 []int
	slice4 = append(slice4, 4)
	fmt.Println(slice1, slice2, slice3, slice4)

}

func mapFn() {
	//go中map的key的类型必须是可以通过==进行比较的，所以slice和浮点型不适合作为map的key，map的value类型没有限制

}

// 就地反转一个slice
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 尽管slice是引用类型，但使用append函数为slice追加元素时仍然需要重新接收结果
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//slice仍有增长空间，扩展slice容量
		z = x[:zlen]
	} else {
		//slice已无空间，为它分配一个新的底层数组
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
