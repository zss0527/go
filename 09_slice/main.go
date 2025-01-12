package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func main() {
	/*
		declare and init way is same with Array
			var variable []type
			var variable = []type{v1,v2}
			variable := []type{v1,v2}

		create slice from existed Array or Slice
			variable := arr[a:b]
			variable := slice[a:b]

		create slice via make()
			var variable = make([]type,len,cap)

		append slice via append
		    var variable = append(sliceA,data1,data2,sliceB)
	*/
	var arr1 []int
	// arr1[0] = 5
	fmt.Printf("arr1: %v,\n type: %T,\n length: %v\n", arr1, unsafe.Sizeof(arr1), len(arr1))
	fmt.Println(arr1 == nil)

	var arr2 = []int{1, 2, 3}
	fmt.Printf("arr2: %v,\n type: %T,\n length: %v\n", arr2, unsafe.Sizeof(arr2), len(arr2))
	arr2[2] = 4
	fmt.Printf("arr2: %v,\n type: %T,\n length: %v\n", arr2, unsafe.Sizeof(arr2), len(arr2))

	arr := [5]int{1, 2, 3, 4, 5}
	b := arr[:] //create slice from existed Array, retrieve all data in the arr

	c := arr[1:4] // retrieve data in arr from index 1 to 4
	d := arr[2:]  // from index 2 to the last one
	fmt.Printf("%v------%T\n", arr, arr)
	fmt.Printf("%v------%T\n", b, b)
	fmt.Printf("%v------%T\n", c, c)
	fmt.Printf("%v------%T\n", d, d)

	e := []string{"c", "java", "go", "js/ts", "python"}
	f := e[2:] // create slice from existed slice
	fmt.Printf("%v------%T\n", f, f)
	fmt.Printf("e-data: %v, length: %d  capability: %d\n", e, len(e), cap(e))
	fmt.Printf("f-data: %v, length: %d  capability: %d\n", f, len(f), cap(f))
	g := e[1:3]
	h := e[:3]
	fmt.Printf("g-data: %v, length: %d  capability: %d\n", g, len(g), cap(g))
	fmt.Printf("h-data: %v, length: %d  capability: %d\n", h, len(h), cap(h))

	sliceA := make([]int, 4, 7)
	fmt.Println(sliceA, len(sliceA), cap(sliceA))
	sliceA[1] = 12
	sliceA[3] = 22
	// sliceA[4] = 33
	fmt.Println(sliceA)

	//can't expand slice via index, should use append(slice,data) method
	sliceA = append(sliceA, 33, 34)
	fmt.Println(len(sliceA), cap(sliceA), sliceA)

	//if capability exceed previous value, the capability will be doubled(not very accurate)
	sliceA = append(sliceA, 45, 46)
	fmt.Println(len(sliceA), cap(sliceA), sliceA)

	sliceB := []int{100}
	//append another slice at the end
	sliceC := append(sliceA, sliceB...)
	fmt.Println(sliceC)

	//slice is reference type
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	s2[0] = 8
	fmt.Println(s1, s2)

	//copy() function copy slice
	s3 := make([]int, 4, 8)
	copy(s3, s1)
	s3[0] = 22
	fmt.Println(s1, s3)

	//no delete function for slice, can use append to achive delete
	//delete 3rd data
	s3 = append(s3[:2], s3[3:]...)
	fmt.Println(s3)

	//change string value via slice
	str1 := "你好 golang"
	runeStr := []rune(str1)
	fmt.Println(runeStr)
	runeStr[0] = '我'
	fmt.Println(string(runeStr))

	//Selection sort and bubble sort
	intList := []int{2, 5, 1, 3, 7, 0, 6}
	sort.Ints(intList)
	fmt.Println(intList)
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)

}
