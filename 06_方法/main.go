package main

import (
	"fmt"
	"image/color"
	"math"
)

// 方法声明
// 和普通函数声明类似，只是在函数名字前面多了一个参数，这个参数把这个方法绑定到这个参数对应的类型上
// 这个类型可以是基本类型也可以是复合类型（但不能是指针或者接口类型），这个参数叫做方法的接收者
type Point struct{ X, Y float64 }

// 包级别的函数
func Distance(p, q Point) float64 { //普通方法
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 类型Point的方法，方法名字是Point.Distance
func (p Point) Distance(q Point) float64 { //Point类型的方法
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

// Path.Distance在内部使用了Point.Distance来计算线段相邻两点之间的距离
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// 指针接收者,方法名字是(*Point).ScaleBy，必须带上小括号，否则会被解析成*(Point.ScaleBy)
// 实际项目中一般遵循，如果Point的任何一个方法使用指针接收者，那么所有的Point方法都应该使用指针接收者
// 命名类型Point与指向它们的指针*Point是唯一可以出现在接收者声明处的类型，不允许本身是指针的类型进行方法声明
/*
type P *int
func (P) f(){}//不可以
*/
/*
三种方式：
1.
	r := &Point{1,2}
	r.ScaleBy(2)
2.
	p := Point{1,2}
	pptr := &p
	pptr.ScaleBy(2)
3.
	p := Point{1,2}
	(&p).ScaleBy(2)

这种方式是不行的,因为没有声明变量，临时变量没法获取地址
&(Point{1,2}).ScaleBy(2)

*/
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	pptr := &p
	fmt.Println(Distance(p, q)) //函数调用
	pptr.Distance(q)

	/*
		p.Distance叫做selector，选择方法和字段是同样的原理
		由于方法和字段来自于同一个命名空间，因此在Point结构类型中声明同样的方法和字段会冲突
		每一个类型有它自己的命名空间，所以可以在不同的类型中使用名字Distance作为方法名（只是接收参数者不同而已
		即类型拥有的所有方法名必须是唯一的，但不同类型可以使用相同的方法名
	*/
	fmt.Println(p.Distance(q)) //方法调用

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	//如果接收者p是Point类型的变量，但方法要求一个*Point接收者，我们可以简写成
	p.ScaleBy(2)
	//实际编译器会对变量进行&p的隐式转换。
	//只有声明变量才允许这样做，其它临时变量例如：结构体字段p.X、数组或者slice的元素priem[3]，不能够对一个不能去地址的Point接收者参数调用*Point方法，因为无法获取临时变量的地址
	// Point{1,2}.ScaleBy(2)//编译错误

	/*
		对于指针接收者方法的三种调用情形
		1. 实参接收者和形参接收者是同一个类型，比如都是T或者*T
			Point{1, 2}.Distance(q)
			pptr.ScaleBy(2)
		2. 实参接收者是T类型的变量而形参接收者是*T类型，编译器会隐式地获取变量的地址
			p.ScaleBy(2)  //隐式转换为(&p).ScaleBy(2)
		3. 实参接收者是*T类型而形参接收者是T类型，编译器会隐式地接引用接收者，获取实际的值
			pptr.Distance(q)  //隐式转换为(*pptr).Distance(q)
	*/

	//能够通过类型为ColoredPoint的接收者调用内嵌类型Point的方法
	var cp ColoredPoint
	cp.X = 1
	cp.Point.Y = 2
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var pp = ColoredPoint{Point{1, 1}, red}
	var qq = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(pp.Distance(qq.Point)) // "5"
	pp.ScaleBy(2)
	qq.ScaleBy(2)
	fmt.Println(pp.Distance(qq.Point)) // "10"
	//go语言中通过结构体嵌套使外层结构体变量可以直接调用内层结构体方法和字段
	//但是并不是因为外层继承了内层，而是编译器为外层结构体提供了包装器方法，而且从pp.Distance(qq.Point)也可以看出，Distance必须接受qq.Point,不可以直接接受qq
	//编译器提供的包装器方法类似于
	/*
		func (p ColoredPoint) Distance(q Point) float64 {
		    return p.Point.Distance(q)
		}
		func (p *ColoredPoint) ScaleBy(factor float64) {
		    p.Point.ScaleBy(factor)
		}
	*/

	//go语言中的封装是通过结构体字段方法首字母大小写实现的，而且封装的级别是包而不是类型

}

/*
结构体内嵌组成类型
*/
type ColoredPoint struct {
	Point
	Color color.RGBA
}
