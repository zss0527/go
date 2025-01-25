package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
go中的复合数据类型是由基本数据类型以各种方式组合而构成的，主要包括数组、slice、map和struct
其中数组和结构体是聚合类型，它们的值由内存中的一组变量构成，数组的元素具有相同的类型，结构体的元素数据类型可以不同；
数组和结构体长度固定，slice和map是动态数据结构，长度可以动态增长；
*/
type Employee struct {
	ID                int
	Name              string
	Adress, Position  string
	Salary, ManagerID int
}

type Point struct{ x, y int }

func main() {
	// arrayFn()
	// sliceFn()
	// mapFn()
	// structFn()
	JsonFn()

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
	//可以通过make创建一个map并分配内存
	ages1 := make(map[string]int)
	//也可以通过map字面量创建一个带初始化键值的map
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	//通过下标访问map元素
	ages1["larry"] = 29
	/*
		但是向一个零值map（未分配内存）中设置元素会导致宕机
	*/
	var ages3 map[string]int
	// ages3["pitiao"] = 23

	//可以通过内置的delete从map中删除元素
	delete(ages3, "alice")
	//delete操作总是安全的，即使所要删除的键并不存在

	//map元素的增长可能导致已有元素再散列，map里的元素不是变量，所以无法获取map元素的地址
	// _ = &ages2["charlie"]

	//对map操作并根据该操作的成功与否做判断的常用结构
	if age, ok := ages2["bob"]; !ok {
		fmt.Println(age)
	}

	//go没有提供Set类型,在go中常用map的key模拟Set
	dup()
}

func structFn() {
	/*
		结构体是将零个或者多个任意类型的命名变量组合在一起的聚合数据类型，每个变量叫做结构体的成员
	*/

	//定义结构体
	/*
		type Employee struct {
			ID                int
			Name              string
			Adress, Position  string
			Salary, ManagerID int
		}
	*/
	//结构体成员变量通常一行写一个，变量名在前类型在后，多行之间不用","，相同类型的连续成员变量可以写在一行
	//成员变量的顺序是结构体类型的一部份，就像容量和长度是某个slice类型的一部份一样
	//首字母大写的结构体成员变量是可以导出的，一个结构体中可以同时包含可导出和不可导出成员变量

	//声明某种结构体变量
	//方式1: 直接声明一个零值的值类型结构体变量
	var dilbert Employee //此时该结构体变量已分配内存，变量是零值，内部成员变量也是对应类型零值
	//方式2: 通过new声明一个零值的指针类型结构体变量
	_ = new(Employee) //用new函数可以声明指针类型的结构体变量
	//方式3: 通过结构体字面量声明并初始化一个值类型的结构体变量
	_ = Point{1, 2}
	//方式4: 通过取结构体字面量地址的方式声明并初始化一个指针类型的结构体变量
	_ = &Point{3, 4}

	//结构体成员都是变量，可以通过点号访问结构体成员变量，可以获取结构体变量和结构体成员变量地址
	dilbert.Salary += 500
	position := &dilbert.Position
	*position = "Senior " + *position
	//结构体中，点号既可以用在结构体值变量上也可以用在结构体地址变量上，go内部自动会转变
	var employeeOfTheMOnth *Employee = &dilbert
	employeeOfTheMOnth.Position += " (proactive team player)"
	(*employeeOfTheMOnth).Position += " (proactive team player)" //go内部会将上面的语句自动转为这种

	//对于返回Employee结构体指针的函数可以直接链式调用
	EmployeeByID(3).Position = "ddddd"
	//如果返回的是Employee左边的表达式无法识别出是一个变量
	// EmployeeByID1(3).Position = "ddddd"

	//命名结构体s中不能包含相同结构体类型s的成员变量，但是可以包含相同结构体类型s的指针成员变量
	type tree struct {
		value       int
		left, right *tree
	}

	//结构体初始化有两种方式1:结构体字面量（要求顺序保持和创建该结构体时的保持一致）， 2:通过指定部份或者全部成员变量的名称和值来初始化；
	//两种方式不可以同时使用
	_ = Point{1, 2} //结构体字面量
	_ = Point{x: 1} //通过指定成员变量初始化，没有被初始化的成员变量值为类型零值

	//结构体值可以直接作为函数参数或返回值，值传递不会影响外面原来的结构体
	//结构体指针也可以作为函数参数或者返回值，引用传递函数内部的修改会直接影响外面原来的结构体，对于复杂结构体指针传递效率更高

	//如果一个结构体所有成员变量都是可比较的（不含slice），那么这个结构体就是可比较的，比较的方式是按照成员变量顺序挨个比较
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p == q) //false

	/*
		结构体嵌套和匿名成员
	*/
	type Circle1 struct{ X, Y, Radius int }
	type Wheel1 struct{ X, Y, Radius, Spokes int }
	var w1 Wheel1
	w1.X = 8
	w1.Y = 9
	w1.Radius = 5
	w1.Spokes = 20
	//结构体嵌套重构
	type Point struct{ X, Y int }
	type Circle2 struct {
		Center2 Point
		Radius  int
	}
	type Wheel2 struct {
		Circle2 Circle2
		Spokes  int
	}
	var w2 Wheel2
	w2.Circle2.Center2.X = 8
	w2.Circle2.Center2.Y = 9
	w2.Circle2.Radius = 5
	w2.Spokes = 20
	//结构体匿名嵌套重构：结构体内部可以定义不带名称的结构体成员，只需要指定类型即可，这种结构体成员称为匿名成员，匿名成员的名称和类型相同
	//这个你名成员必须是一个命名类型或者指向命名类型的指针
	type Circle struct {
		Point
		Radius int
	}
	type Wheel struct {
		Circle
		Spokes int
	}
	var w Wheel
	w.X = 8
	w.Y = 9 //等价于w.Circle.Point.X = 8
	w.Radius = 8
	w.Spokes = 20
	//结构体字面量并没有快捷方式来初始化结构体
	// w = Wheel{8,9,10,20}//编译错误
	// w = Wheel{X:8,Y:9,Radius: 5, Spokes: 20}//编译错误
	w = Wheel{Circle{Point{8, 9}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{8, 9},
			Radius: 5,
		},
		Spokes: 20, //结尾必须有逗号
	}
	//不能在同一个结构体定义两个相同类型的你名成员
	//外围的结构体类型获取的不仅是你名成员的内部变量，还有相关的方法，Go是通过这种组合实现面向对象编程方式的

}

func JsonFn() {
	/*
	   go通过标准库encoding/json,encoding.xml分别json和xml格式的编码和解码提供支持
	*/
	type Movie struct {
		Title  string
		Year   int  `json:"released"` //称为标签，用来实现成员变量和json字段名称的同步
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	//go的数据结构转换为json称为marshal，通过json.Marshal来实现
	data1, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data1)

	//json.MarshalIndent用来格式化marshal后的json格式
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//将json字符串转换为go的响应数据结构称为unmarshal，通过json.Unmarshal来实现
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

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

func dup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {

			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func EmployeeByID(id int) *Employee {
	return &Employee{}
}

func EmployeeByID1(id int) Employee {
	return Employee{}
}
