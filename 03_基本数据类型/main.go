package main

/*
go的数据类型分为4大类型：基础类型（basic type）、聚合类型（aggregate type）、引用类型（reference type）和接口类型（interface type）
基础类型：数字（number）、字符串（string）和布尔型（boolean）
聚合类型：数组（array）和结构体（struct）
引用类型：指针（pointer）、切片（slice）、映射（map）、函数（func）和通道（channel）
接口类型： interface

3.1 整数
	特定长度的有符号和无符号整型：int8,int16,int32,int64,uint8,uint16,uint32,uint64
	平台相关长度的整型：int和uint
	rune是int32类型的同义词，常常用于指明一个值是Unicode码点
	byte是uint8类型的同义词，强调一个值是原始数据，而非量值
	uintptr大小不明确的无符号整数，但足以完整存放指针，常用于底层编程
3.2 浮点数
	float32和float64，遵从IEEE754标准
3.3 复数
	complex64和complex128，内置的real和imag函数用来提取复数的实部和虚部
	var x complex128 = complex(1,2)//1+2i
	var y complex128 = complex(3,4)//3+4i
	real(x*y)//-5
	imag(x*y)//10
3.4 布尔值
	true，false
3.5 字符串
	go中字符串是不可变的字节序列，习惯上文本字符串被解读成UTF-8编码的Unicode码点序列
	len(str)返回一个字符串长度
	str[i:j]子串生成操作 str[:j] str[i:]
	str1+str2两个字符串连接操作，字符串不能直接和另一个非字符串通过+连接
	字符串内部数据不可以修改str[i]='L'错误
	3.5.1 字符串字面量
		带双引号的字节序列字面量"Hello Go"，这种形式的字符串内以\开始的转义序列会生效；
		原生的字符串字面量`Hello Go`，这种形式的字符串内转义序列不生效；
	3.5.2 Unicode
		ASCII用7位表示128个字符，Unicode用int32（rune）
	3.5.3 UTF-8
		UTF-8以字节为单位对Unicode码点作变长编码，UFT-8是现行的一种Unicode标准；
	3.5.4 字符串和字节slice
		4个操作字符串的包：、
			strings：搜索、替换、比较、修整、切分和字符串连接
			bytes：主要用于操作字节slice []byte,bytes.Buffer效率更高
			strconv：转换bool、int、float为对应的字符串，或者反过来
			unicode：判别文字符号值特性的函数，IdDigit、IsLetter、IsUpper、IsLower
	3.5.5 字符串和数字的相互转换
		方法一：fmt.Sprintf("%d",x)
		方法二：strconv.Itoa(str)
3.6 常量
	常量是一种表达式，可以保证在编译阶段就计算出表达式的值 const pi = 3.14
	3.6.1 常量生成器iota
		常量声明中，iota从0开始取值，逐项加1
		const(
			Sunday int = iota
			Monday
			Tuesday
			Wednesday
			Thursday
			Friday
			Saturday
		)
		Sunday值为0，Monday为1，以此类推
	3.6.2 无类型常量
		算术精度高于原声的机器精度，至少精度达到256位


*/
