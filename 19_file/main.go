package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
 1. read file
    1.1
    file, err := os.Open("path")
    file.Read()
    defer file.Close()
    1.2
    file, err := os.Open("path")
    reader := bufio.NewReader(file)
    line, err := reader.ReadString('\n')
    defer file.Close()
    1.3
    ioutil.ReadFile("path")

 2. write file
    2.1
    file, err := os.OpenFile("path",os.O_CREATE|os.O_RDWR,0666)
    file.Write([]byte(str))
    file.WriteString("str")
    file.Close()
    2.2
    file, err := os.OpenFile("path",os.O_CREATE|os.O_RDWR,0666)
    writer := bufio.NewWriter(file)
    writer.WriteString("str")
    writer.Flush()
    file.Close()
    2.3
    err := ioutil.WriteFile("path",[]byte(str),0666)
*/
func main() {
	// read1()
	// read2()
	// read3()

	// write1()
	// write2()
	write3()

}

func read1() {
	var tempSlice = make([]byte, 128)
	var contentSlice []byte
	file, err := os.Open("./readme.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("open file failed.")
		return
	} else {
		fmt.Println("open file success!")
		//1.file.Read()
		for {
			count, err := file.Read(tempSlice)
			if err == io.EOF { //means read finished
				break
			}
			if err != nil {
				fmt.Println("read file failed!")
				return
			} else {
				fmt.Println(count, "bytes")
				contentSlice = append(contentSlice, tempSlice[:count]...)
			}
		}
		fmt.Println(string(contentSlice))
	}
}

func read2() {
	file, err := os.Open("./readme.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("open file failed.")
		return
	} else {
		fmt.Println("open file success!")
		reader := bufio.NewReader(file)
		var fileStr string
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				fileStr += line
				break
			}
			if err != nil {
				fmt.Println("read line failed!")
				return
			}
			fileStr += line
		}
		fmt.Println(fileStr)
	}
}

func read3() {
	byteStr, err := ioutil.ReadFile("./readme.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}

func write1() {
	file, err := os.OpenFile("./writeme.txt", os.O_RDWR|os.O_TRUNC|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		file.WriteString(fmt.Sprintf("dsdsd放德弗里斯的见风使舵了dsdfsdfsd的粉丝地方%v\n", i))
	}
	fmt.Println("write sucessfully...")
}

func write2() {
	file, err := os.OpenFile("./writeme.txt", os.O_RDWR|os.O_TRUNC|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < 6; i++ {
		writer.WriteString(fmt.Sprintf("dsdsd放德弗里斯的见风使舵了dsdfsdfsd的粉丝地方%v\n", i))
	}
	writer.Flush()
	fmt.Println("write success!")
}

func write3() {
	err := ioutil.WriteFile("./writeme.txt", []byte("dflsdfjsdlfsdjljddd44444"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("write success.")
}
