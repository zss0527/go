package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

/*
接口类型是对其它类型行为的概括与抽象
通过使用接口我们可以写出更加灵活和通用的函数，这些函数不用绑定在一个特定的类型实现上
Go语言的接口相对于java是隐式实现的：
	对于一个具体的类型，无需声明它实现了哪些接口，只要提供接口所需的方法即可；
	这种设计让我们无需改变已有类型的实现，就可以为这些类型创建新的接口，对于不能修改的类很有用

具体类型：
	具体类型指定了它所含数据的精确布局，还暴露了基于这个精确布局的内部操作；
	如果知道了一个具体类型的数据，那么就精确知道了它是什么以及它能干什么；
抽象类型：
	接口是一种抽象类型，它并没有暴露所含数据的布局或者内部结构，当然也没那些数据的基本操作，它所提供的仅仅是一些方法；
	如果知道了一个接口类型的数据，你将无从知道它是什么，您能知道的仅仅是它能做什么（提供了哪些方法）；
*/

func main() {
	/*
		实现接口：如果一个类型实现了一个接口要求的所有方法，那么这个类型就实现了这个接口；
			*os.File类型实现了io.Reader，Writer，Closer，和ReadWriter接口；
			*bytes.Buffer实现了Reader，Writer，和ReadWriter这些接口，但是它没有实现Closer接口因为它不具有Close方法；
			Go程序员通常说一个具体类型“是一个”（is-a）特定的接口类型，这代表该具体类型实现了该接口，比如*bytes.Buffer是一个io.Writer，*os.Files是一个io.ReadWriter；
		接口的赋值规则：仅当一个表达式实现了一个接口时，这个表达式才可以赋值给接口（右边表达式也是一个接口时改规则也有效）

		接口封装了所对应的类型和数据，只有通过接口暴露的方法才可以调用，类型的其它方法无法通过接口调用（这点类似java）；

		非空接口类型（方法）通常由一个指针类型来实现（接收者用指针类型）特别是接口中的一些方法暗示会修改接收者的情形；
		但指针类型不是实现接口的唯一类型，引用类型、函数类型甚至基础类型都可以，但通常使用指针类型作为实现方法的接收者；
		一个具体类型可以是像很多不相关的接口；
	*/
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	// w = time.Second//错误

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// rwc = new(bytes.Buffer)//错误

	w = rwc
	// rwc = w//错误

	w.Write([]byte("Hello Go"))
	//尽管w的实际赋值表达式是os.Stdout类型，但是w仍然只能调用io.Writer接口的方法，没法调用io.Stdout的额外方法
	// w.Close()//错误

	/*
		空接口类型interface{}
		空接口类型不包含任何方法，可以把任何值赋给空接口类型，类似于js中的any
		变长参数函数中，可以通过...interface{}接收不限个数的任意类型的参数，但是这样的话就没法使用这些参数的任何值和方法（因为空接口不含任何方法）可通过类型断言解决这个问题
	*/
	var any interface{}
	any = true
	any = 12.24
	any = map[string]int{"one": 1}
	fmt.Println(any)

}

// 一个接口类型定义了一套方法，如果一个具体类型要实现该接口，那么必须实现接口类型定义中的所有方法；
// 接口可以通过嵌入包含被嵌入接口的方法，当然也可以自己重新在接口里声明这些方法；
type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface{ Close() error }
type ReadWriter interface {
	Reader
	Writer
}
type ReadWriterM interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}
type Text interface {
	Pages() int
	Words() int
	PageSize() int
}
type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP3", "WAV"
}
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // e.g., "MP4", "WMV"
	Resolution() (x, y int)
}

// 可以定义一个Streamer接口来代表它们之间相同的部分而不必对已经存在的类型做改变
// 从这个逻辑可以看出，提取共性或者其它原因创建新的接口是随时可以的，且不影响原来类型的定义，这一点和java非常不同
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

// 接口值：概念上一个接口类型的值包含两个部分：一个具体类型和该类型的一个值，二者称为接口的动态类型和动态值；
// go是静态型语言，类型仅仅是一个编译时的概念，所以类型不是一个值，只是用类型描述符来提供每个类型具体信息
func 接口值() {
	//接口的零值是把它的动态类型和动态指都设置为nil
	//一个接口值是否是nil取决于它的动态类型，可以用==或者!=来检测一个接口值是否是nil，调用一个nil接口的任何方法都会错误
	var w io.Writer
	if w == nil {
		fmt.Println("w is nil")
	}

	//把一个*os.File类型的值赋值给了w，把一个具体类型隐式转换为一个接口类型，它与显式的转换io.Writer(os.Stdout)等价
	//接口值的动态类型会设置为指针类型*os.File的类型描述符，它的动态值会设置为os.Stdout的一个副本，即一个指向代表进城的标准输出的os.File类型的指针
	w = os.Stdout
	//调用该接口的Write方法会实际调用(*os.File).Write方法
	//实际上编译器无法知道一个接口值的动态类型会是什么，所以通过接口来做调用会使用动态分发：
	//编译器会生成一段代码来从类型描述符拿到名为Write的方法地址，再间接调用该方法地址，调用的接收者就是接口值的动态值即os.Stdout
	//所以实际等效于 os.Stdout.Write([]byte("Hello Go"))
	w.Write([]byte("Hello Go"))

	//动态类型是*bytes.Buffer,动态值是一个指向新分配缓冲区的指针
	//调用Write方法的机制和上述一样，等价于(*bytes.Buffer).Write方法，方法的接收者是缓冲区的地址
	w = new(bytes.Buffer)
	w.Write([]byte("Hello GoLang"))

	//动态类型和动态值都设置成了nil
	w = nil

	//接口值可以通过==和!=进行比较，如果两个接口值都是nil或者二者的动态类型完全一致且二者动态值相等(==)那么这两个接口相等
	//因为接口值可以进行比较，所以接口可以作为map的key

	names := []string{"aaa", "bbbb"}
	sort.Sort(StringSlice(names))
	//sort包内置了StringSlice类型以及一个直接排序的Strings函数，所以上面的可以简写为
	sort.Strings(names)

	//类型断言：x.(T)
	//类型断言是一个作用在接口上的操作，格式为x.(T),其中x是一个接口类型的表达式，而T是一个类型（称为断言类型）
	//如果断言类型T是一个具体类型，那么类型断言会检查x的动态类型是否就是T，如果检查成功那么类型断言的结果就是x的动态值，如果失败则崩溃
	fg := w.(*os.File) //f == os.Stdout，成功所以fg此时明确为类型为os.Stdout
	fmt.Println(fg)
}

// 使用sort.Interface来排序
// 在很多语言中，排序算法都是和序列数据类型关联，同时排序函数和具体类型元素关联。相比之下，Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。相反，它使用了一个接口类型sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片
//一个排序算法需要知道三个信息：序列长度、比较连个元素的含义以及如何交换两个元素
//sort的Interface接口就三个方发：
/*
type Interface interface {
    Len() int
    Less(i, j int) bool // i, j are indices of sequence elements
    Swap(i, j int)
}
*/
//所以要对序列排序需要先确定一个实现了如上三个方法的类型，接着把sort.Sort函数应用到上面这类方法的实例上
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
