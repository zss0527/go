package main

import "fmt"

/*
每一个变量对应着四个属性：字面量（名字）、内存起始地址、类型、值
其中内存起始地址+类型确定了值的内存空间，这个内存空间里的内容就是值

函数也是如此但又有一些变化，函数名是字面量，函数名的地址不可见，函数名的值是函数体的内存起始地址，类型是函数的签名
通过()进行调用，函数名的赋值（函数体内存起始地址赋值）有个前提是类型必须相同

对于指针，值是某个变量的地址，类型是对应的变量的指针类型，地址是这个指针的内存起始地址

&符号用来取某个变量的地址
*符号用来指定某个地址，作用1:指定后获取该地址里的值，作用2:指定后对该地址里的值重新赋值

指针类型
*int *string .....

指针必须在分配内存后才能使用
make可以用于slice，map和channel的初始化并返回引用类型
new用于类型的内存分配并初始化对应值为0，返回指向该内存的指针

golang中指针不能像c那样进行算术运算

值类型：值是定义的那个字面量
引用类型：值也是定义地那个字面量，引用类型作变量被赋值时不会重新创建内存空间，只会将自己的地址改为右侧赋值变量的地址，go中只有三种引用类型slice、map和channel，没法像c++那样声明任意类型的引用类型
指针：值是指向某个存储单元的地址
*/
func main() {

	var a int = 10
	var po *int = &a
	fmt.Println(po)
	p := &a
	fmt.Printf("name: a, value: %v, type: %T, address: %p\n", a, a, &a)
	fmt.Printf("name: p, value: %v, type: %T, address: %p\n", p, p, &p)
	fmt.Printf("name: test, value: %v, type: %T\n", test, test) //can't get func address in go

	fmt.Println(*p)
	*p = 20
	fmt.Println(a)

	pointerParamFn(&a)
	fmt.Println(a)

	/*
		map is reference type, it's value is address, pointer
		so for map, but allocate memory via make before use it
	*/
	// var userInfo map[string]string
	// userInfo["userName"] = "zhangsan"
	// fmt.Println(userInfo)
	var userInfo = make(map[string]string)
	userInfo["userName"] = "larry"
	fmt.Printf("value: %v,  address: %p, type: %T\n", userInfo, userInfo, userInfo)
	fmt.Println(&userInfo, (*&userInfo)["userName"])
	receiveMapFn(userInfo)

	// var sl []int
	// sl[0] = 1
	var sl = make([]int, 4, 4)
	sl[0] = 1
	fmt.Println(sl)

	// var cc *int
	// cc = 100
	var cc = new(int) // var cc *int, then allocate memory for the cc
	fmt.Println(cc, *cc)

}

func test() int { return 1 }

// param is pointer type
func pointerParamFn(x *int) {
	fmt.Println(x)
	*x = 40
}

func receiveMapFn(x map[string]string) {
	fmt.Println("111111", x, &x)
	fmt.Printf("%p\n", x)
}
