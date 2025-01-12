package main

import (
	"encoding/json"
	"fmt"
)

/*
1. struct definition
	type structType struct {  //this structType is public if the first char is capital, otherwise, it's private
		fields
		methods
	}

2. struct initialize
	2.1 var ins structType  //value type, can be used after declared
	2.2 var ins = new(structType)  //pointer
	2.3 var ins = &structType{}    //pointer
	2.4 var ins = structType{      //value
								field1:value1,
								field2:value2,
							}
	2.5 var ins = &structType{      //pointer
								field1:value1,
								field2:value2,
							}
	2.6 var ins = structType{value1,value2}   //value, all field must be include and order by definition
	2.7 var ins = &structType{value1,value2}  //pointer, all field must be include and order by definition

3. struct methods
	3.1 func (receiveName receiveType) methodName(params)(return types){code block}
	3.2 func (receiveName *receiveType) methodName(params)(return types){code block}

4. anonymous fields in struct, generally only used in struct nesting case
	type Person struct {
		string
		int
	}

5. struct nesting
	type User struct {
		UserName string
		Password string
		age int
		Address Address
	}
	type Address struct {
		Name string
		Phone string
		city string
	}

6. Struct inheritance
	struct inheritance is through struct nesting

7. serialization and deserialization of struct
	encoding/json

*/

// self defined type
type myInt1 int
type myFn func(int, int) int

// type alias
type myFloat = float64

type Person struct {
	Name string
	sex  string
	age  int
}

type Animal struct {
	Name  string   `json:"name"`
	Hobby []string `json:hobby`
	age   int
	infos map[string]string
}

type myInt int

func (t myInt) methodForSelfDefinedType() {
	fmt.Println("ddddd")
}

type User struct {
	UserName string
	Password string
	age      int
	Address  Address
	Order    //anynimous feild
}
type Address struct {
	Name  string
	Phone string
	city  string
}

type Order struct {
	id    string
	price float64
}

type Animal1 struct {
	name string
}

// struct inheritance
type Cat struct {
	Animal1
	*Animal
	Age int
}

type Student struct {
	Id     int
	Gender string
	Name   string
}
type Class struct {
	Title    string
	Students []Student
}

func (animal Animal1) run() {
	fmt.Printf("%v is running\n", animal.name)
}

func (cat Cat) miao() {
	fmt.Printf("%v is maiomaiomaio\n", cat.name)
}
func main() {
	type myInt int

	var a myInt = 10
	var b myInt1 = 20
	var c myFloat = 2.14
	fmt.Printf("%v---%T\n", a, a)
	fmt.Printf("%v---%T\n", b, b)
	fmt.Printf("%v---%T\n", c, c)

	//value
	var person1 Person // value type, can be used after declare
	person1.Name = "Larry"
	person1.sex = "male"
	fmt.Printf("value: %v   type: %T\n", person1, person1)
	fmt.Printf("detials: %#v\n", person1)

	pp1 := person1
	pp1.Name = "xxxx"
	fmt.Println(person1, pp1) //values are not same, means struct is value type, not reference type

	//pointer
	//can also visit struct fiels via . if the type is struct pointer
	//p2.name is also (*p2).name
	var person2 = new(Person)
	person2.Name = "zhangsan"
	person2.age = 11
	(*person2).sex = "male"
	fmt.Printf("person2 value: %v   type: %T\n", person2, person2)
	fmt.Printf("person2 detials: %#v\n", person2)
	fmt.Println(person2.Name, (*person2).Name)

	//pointer
	var person3 = &Person{}
	person3.Name = "lisi"
	(*person3).age = 20
	fmt.Printf("person3 value: %v   type: %T\n", person3, person3)
	fmt.Printf("person3 detials: %#v\n", person3)

	//value
	var person4 = Person{
		Name: "lisi",
		age:  14,
	}
	fmt.Printf("person4 value: %v   type: %T\n", person4, person4)
	fmt.Printf("person4 detials: %#v\n", person4)

	//pointer
	var person5 = &Person{
		Name: "wnagwu",
		age:  14,
	}
	fmt.Printf("person5 value: %v   type: %T\n", person5, person5)
	fmt.Printf("person5 detials: %#v\n", person5)

	var person6 = Person{"zhuliu", "male", 18}
	fmt.Printf("person6 value: %v   type: %T\n", person6, person6)
	fmt.Printf("person6 detials: %#v\n", person6)

	var person7 = &Person{"liuqi", "male", 33}
	fmt.Printf("person7 value: %v   type: %T\n", person7, person7)
	fmt.Printf("person7 detials: %#v\n", person7)

	//call struct methods
	var pm = Person{"zhangsan", "male", 22}
	fmt.Println(pm)
	fmt.Println(pm.printInfo("sansan"))
	fmt.Println(pm)
	pm.printInfoP("sisi")
	fmt.Println(pm)
	pp := pm.printInfoA("wuwu")
	fmt.Println(pp)

	// var mi myInt = 5
	// mi.methodForSelfDefinedType()

	var animal Animal
	animal.Name = "pitiao"
	animal.age = 2
	animal.Hobby = make([]string, 6, 6)
	animal.Hobby[0] = "eating"
	animal.Hobby[1] = "sleeping"
	animal.infos = make(map[string]string)
	animal.infos["type"] = "cat"
	animal.infos["isClever"] = "yes"

	fmt.Println(animal)

	var user User
	// var address Address
	user.UserName = "xiaowang"
	user.age = 25
	user.Password = "123"
	user.Address.Name = "addressname"
	user.Address.Phone = "1234567"
	user.Address.city = "sh"
	user.Order.id = "1111"
	user.Order.price = 12.345
	user.id = "2222" //look for this field from the outside to inside
	fmt.Printf("%#v\n", user)

	//struct inheritance
	var cat = Cat{
		Age: 2,
		Animal1: Animal1{
			name: "cat",
		},
		Animal: &Animal{
			Hobby: []string{"eat", "sleeping", "maiomiaomiao"},
		},
	}
	cat.run()
	cat.miao()

	//serialization
	var ani = Animal{
		Name:  "pitiao",
		Hobby: []string{"eat", "sleep", "miaomiaomaio"},
		age:   2, //private fields will not be visited by json tools
		infos: map[string]string{"isClever": "no"},
	}
	fmt.Printf("%#v\n", ani)
	jsonByte, _ := json.Marshal(ani)
	jsonStr := string(jsonByte)
	fmt.Println("jsonStr:", jsonStr)

	//deserialization
	var jstr = `{"Name":"pitiao","Hobby":["eat","sleep","miaomiaomaio"]}`
	var ani1 Animal
	err := json.Unmarshal([]byte(jstr), &ani1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", ani1)

	class := Class{
		Title:    "02",
		Students: make([]Student, 0),
	}
	for i := 1; i < 10; i++ {
		s := Student{
			Id:     i,
			Gender: "male",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		class.Students = append(class.Students, s)
	}
	fmt.Println(class)
	jsonB, _ := json.Marshal(class)
	fmt.Println(string(jsonB))

	j := `{"Title":"02","Students":[{"Id":1,"Gender":"male","Name":"stu_1"},{"Id":2,"Gender":"male","Name":"stu_2"},{"Id":3,"Gender":"male","Name":"stu_3"},{"Id":4,"Gender":"male","Name":"stu_4"},{"Id":5,"Gender":"male","Name":"stu_5"},{"Id":6,"Gender":"male","Name":"stu_6"},{"Id":7,"Gender":"male","Name":"stu_7"},{"Id":8,"Gender":"male","Name":"stu_8"},{"Id":9,"Gender":"male","Name":"stu_9"}]}`
	var cc = &Class{}
	json.Unmarshal([]byte(j), cc)
	fmt.Printf("%#v\n", cc)

}

// struct methods
func (p Person) printInfo(name string) (int, string) {
	fmt.Println(p.Name)
	p.Name = name
	return p.age, p.Name
}

func (p *Person) printInfoP(name string) (int, string) {
	fmt.Println(p.Name)
	p.Name = name
	return p.age, p.Name
}

func (p Person) printInfoA(name string) Person {
	p.Name = name
	return p
}
