package main

import "fmt"

/*
1. interface definition
	type interfaceName interface {
		method1(params list) return list
		method2(params list) return list
	}
	generally, interface is implemented by struct

2. blank interface
	have no any fields and methods, used to abstract common features and create this obj
	blank interface can be used params of function, meaning the function can receive any type params
	blank interface unsupport index.

3. type assert
	interface value is constructed by a certain type and the certain type value, these two parts named dynamic type and dynamic value
	if we watn to know the blank interface type, we can use type assert.
	interfaceTypeVariable.(T)
		x:interface type variable
		T:possible types

4. struct value receiver, struct pointer receiver
	if receiver is value type, struct value type and pointer type variable can be received.
	if receiver is pointer type, only struct pointer type variable can ve received.

5. multiple interface implement

6. interface nesting
	type interface1 interface {method1()}
	type interface2 interface {method2()}
	type interface3 interface {interface1 interface2}

*/

type Device interface{}

type Usber interface {
	start(string, string) string
	stop() bool
}

type Phone struct {
	Name string
}

type Camera struct {
}

// Phone must implement all methods if Phone want to implement Usb interface
func (p Phone) start(who, why string) string { //value receiver
	fmt.Println(who, "started", p.Name, "because", why)
	return "started successfully!"
}
func (p Phone) stop() bool { //value receiver
	fmt.Println(p.Name, "stoped...")
	return true
}

func (p *Camera) start(who, why string) string { //value receiver
	fmt.Println(who, "started the camera", "because", why)
	return "started successfully!"
}
func (p *Camera) stop() bool { //pointer receiver
	fmt.Println("camera stoped...")
	return true
}

type Computer struct {
}

func (c Computer) work(usb Usber) {
	if _, ok := usb.(Phone); ok {
		usb.start("111", "222")
	} else {
		usb.stop()
	}
}

func main1() {
	myPhone := Phone{
		Name: "iphone13",
	}
	myPhonePointer := &Phone{
		Name: "iphone14",
	}

	myPhone.start("zss0527", "want to play wangzherongrao!")

	//value receiver, both value variable and pointer variable can be received.
	var myPhoneWithUsb Usber = myPhone //implement
	myPhoneWithUsb.start("zss0527", "want to listen to music")
	var myPhoneWithUsbPointer Usber = myPhonePointer //implement
	myPhoneWithUsbPointer.start("dddd", "tddddd")

	//pointer receiver, only pointer can be received
	// var camera Usber = Camera{}  //not work
	var camera Usber = &Camera{}
	camera.start("pitiao", "want to take photo")
	camera.stop()

	computer := Computer{}
	computer.work(myPhone)
	computer.work(camera)

	//blank interface example1
	var device Device
	var str = "hello golang"
	device = str
	fmt.Printf("value: %v, type: %T\n", device, device)
	var num = 20
	device = num
	fmt.Printf("value: %v, type: %T\n", device, device)
	//blank interface example2
	var a interface{}
	a = 20
	v, ok := a.(int)
	//type assert
	if ok {
		fmt.Println("v:", v)
	}
	a = "hello"
	a = false
	fmt.Printf("value: %#v\n", a)
	//blank interface example3
	show(34)
	show(true)
	show([]int{1, 2, 3})
	show([]interface{}{1, "dddd", false, make(map[string]interface{})})

	PrintEveryTypeFn("hello")
}

// receive blank interface as param
func show(in interface{}) {
	fmt.Printf("%#v\n", in)
}

func PrintEveryTypeFn(x interface{}) {
	// if _, ok := x.(string); ok {
	// 	fmt.Println("string type")
	// }
	switch x.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case bool:
		fmt.Println("bool type")
	default:
		fmt.Println("not found this type")
	}
}

type Animal interface {
	SetName(string)
	GetName() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

func (d *Cat) SetName(name string) {
	d.Name = name
}

func (d Cat) GetName() string {
	return d.Name
}

func main() {
	var d Animal = &Dog{
		Name: "xiaohei",
	}
	fmt.Printf("d.GetName(): %v\n", d.GetName())
	d.SetName("aqi")
	fmt.Println(d)

	var cat Animal = &Cat{
		Name: "pitiao",
	}
	fmt.Printf("cat.GetName(): %v\n", cat.GetName())
	cat.SetName("pitiao1")
	fmt.Println(cat)

	//multiple interface implemention
	dog := &Dog{
		Name: "xiaohuang",
	}
	var dog1 Animal1 = dog
	var dog2 Animal2 = dog
	fmt.Printf("dog2.GetName(): %v\n", dog2.GetName())
	dog1.SetName("xiaozi")
	fmt.Println(dog)

	//interface nesting
	var dog3 Animal3 = dog //all methods must be implemented in Dog
	fmt.Printf("dog.GetName(): %v\n", dog.GetName())
	dog3.SetName("dahuang")
	fmt.Println(dog3)

	userInfo := make(map[string]interface{})
	userInfo["userNmae"] = "zhangsan"
	userInfo["age"] = 20
	userInfo["hobby"] = []string{"paly", "sleep"}

	fmt.Printf("userInfo[\"age\"]: %v\n", userInfo["age"])
	//blank interface unsupport index
	//userinfo["hobby"][1]   //don't work
	v1, _ := userInfo["hobby"].([]string)
	fmt.Printf("v[2]: %v\n", v1[1])

	var address = Address{
		Name:  "lisi",
		Phone: "1234567",
	}
	userInfo["address"] = address
	//userInfo["address"].Phone  //don't work
	v2, _ := userInfo["address"].(Address)
	fmt.Printf("v2.Name: %v\n", v2.Name)

}

type Animal1 interface {
	SetName(string)
}

type Animal2 interface {
	GetName() string
}

// interface nesting
type Animal3 interface {
	Animal1
	Animal2
}

type Address struct {
	Name  string
	Phone string
}
