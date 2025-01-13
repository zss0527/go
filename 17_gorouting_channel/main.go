package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
1. gorouting
	gorouting is a user-level thread and will not cause system-level thread switching.

2. channel
	channel is one reference type in go
	differents gorouting commucate via channel
	2.1 channel declare
		var ch1 chan int
		var ch2 chan string
	2.2 channel init
		ch1 = make(chan int 10)
	2.3 channel operations
		sent infos to channel
			ch1 <- 10
		receive infos from channel
			x := <- ch1
		close channel
3. one way channel, two ways channel can be assigned to one way channel, but not the other way around.
	var ch1 chan<- int
	var ch2 <-chan int
	ch1 = make(chan<- int,2)
	ch2 = make(<-chan int,2)

4. select multiplexing
	select {
		case <- ch1:
			...
		case data := <- ch2
			...
		case ch3<-data:
			...
		default:
			...
	}

5. Mutex lock, RWMutex lock


*/

var wg sync.WaitGroup

func main1() {

	go test1Fn() //start a gorouting
	for i := 0; i < 10; i++ {
		fmt.Println("你好 Golang -", i)
		time.Sleep(time.Duration(time.Millisecond * 50))
	}
	//proactively block the main thread and wait for the execution of the goroutine to complete
	//time.Sleep(time.Second)
	wg.Wait() // main thread will exit after gorouting counter is 0
	fmt.Println("main thread exited!")
}

func test1Fn() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello Golang -", i)
		time.Sleep(time.Duration(time.Millisecond * 100))
	}
	wg.Done() //gorouting counter sub 1
}

func test2Fn() {
	wg.Add(1) // gorouting counter add 1
	for i := 0; i < 10; i++ {
		fmt.Println("grouting test -", i)
		time.Sleep(time.Duration(time.Millisecond * 100))
	}
	wg.Done() //gorouting counter sub 1
}

func someFuncsInGoForThreadSettings() {
	wg.Add(1)                  // gorouting counter add 1, meaning add lock
	cpuNum := runtime.NumCPU() //get cpu num on this computer
	fmt.Println("cpuNum:", cpuNum)

	runtime.GOMAXPROCS(cpuNum - 1) //set cpu number used in this application
	wg.Done()                      //release lock
}

func test(num int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("gorouting(%v), %v\n", num, i)
		time.Sleep(time.Millisecond * 100)
	}
	defer wg.Done()
}

func main2() {
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go test(i)
	}

	wg.Wait()
	fmt.Println("main thread exited!")
}

func main3() {
	startTime := time.Now().UnixMilli()
	for i := 2; i < 100000; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Print(i, " ")
		}
	}
	endTime := time.Now().UnixMilli()
	fmt.Println("took time:", endTime-startTime)
}

func main7() {
	startTime := time.Now().UnixMilli()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go calculateFn(i)
	}
	wg.Wait()
	endTime := time.Now().UnixMilli()
	fmt.Println("took time:", endTime-startTime)

}

func calculateFn(n int) {
	for i := (n-1)*30000 + 1; i < n*30000; i++ {
		if n > 1 {
			var flag = true
			for j := 2; j < i; j++ {
				if i%j == 0 {
					flag = false
					break
				}
			}
			if flag {
				// fmt.Print(i, " ")
			}
		}
	}
	wg.Done()
}

func main5() {
	//create channel
	ch := make(chan int, 5)

	//sent infos to channel
	ch <- 10
	ch <- 32
	ch <- 28
	fmt.Printf("%#v, %v, %v\n", ch, cap(ch), len(ch))

	//get infos from channel
	x := <-ch
	<-ch
	y := <-ch
	fmt.Println(x, y)

	//channel
	ch1 := make(chan int, 10)
	ch1 = ch
	ch1 <- 31
	ch1 <- 54
	z := <-ch
	m := <-ch1
	fmt.Println(z, m)

	ch6 := make(chan int, 1)
	ch6 <- 2
	//deadlock
	// ch6 <- 8

	ch7 := make(chan string, 2)
	ch7 <- "data1"
	ch7 <- "data2"
	m1 := <-ch7
	m2 := <-ch7
	fmt.Println(m1, m2)
	//deadlock
	// m3 := <-ch7
	// fmt.Println(m3)

	var ch8 = make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch8 <- i + 1
	}
	/*
		must close it before for rang the channel,
		because for range don't know when to finish the for loop,
		so, need producer tell the for range loop explicitly via close(ch).
	*/
	close(ch8)

	for v := range ch8 {
		fmt.Println(v)
	}

}

func main6() {
	var ch = make(chan int, 10)
	wg.Add(1)
	go producer(ch)
	wg.Add(1)
	go consumer(ch)

	wg.Wait()
	fmt.Println("existed...")
}

func producer(ch chan int) {
	for i := 1; i < 10; i++ {
		ch <- i
		fmt.Printf("produce data: %v\n", i)
		time.Sleep(time.Millisecond * 3000)
	}
	close(ch)
	wg.Done()
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Printf("consume data: %v\n", v)
	}
	wg.Done()
}

func main8() {
	startTime := time.Now().UnixMilli()
	n, num := 3, 120000
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16)

	wg.Add(1)
	go putNum(intChan, num)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0; i < n; i++ {
			<-exitChan
		}
		close(primeChan)

		wg.Done()
	}()

	wg.Wait()
	endTime := time.Now().UnixMilli()
	fmt.Println("took time: ", endTime-startTime, "milliseconds")
	fmt.Println("main thread exited..")

}

// put 1-120000 into intChan
func putNum(intChan chan int, num int) {
	for i := 2; i < num; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

// calculate prime
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for v := range intChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- v
		}
	}
	// close(primeChan)
	exitChan <- true
	wg.Done()
}

// printPrime
func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg.Done()
}

func main9() {
	var ch1 chan<- int
	var ch2 <-chan int
	ch1 = make(chan<- int, 3)
	ch2 = make(<-chan int, 3)
	ch3 := make(chan int, 10)
	ch3 <- 44
	ch3 <- 19
	ch1 = ch3
	ch2 = ch3

	fmt.Println(<-ch2)
	ch1 <- 8
	fmt.Println(<-ch3)

	intChan := make(chan int, 10)
	stringChan := make(chan string, 20)

	for i := 0; i < 10; i++ {
		intChan <- i + 1
	}
	for i := 0; i < 10; i++ {
		stringChan <- fmt.Sprintf("str-%d", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Println(v)
		case s := <-stringChan:
			fmt.Println(s)
		default:
			fmt.Println("all data retrived from intChan and stringChan!")
			return
		}
	}

}

func main10() {
	wg.Add(2)
	go sayHello()
	go myTest()

	wg.Wait()
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello, world", i)
	}
	wg.Done()
}

func myTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("errors appeared in myTest():", err)
		}
		wg.Done()
	}()
	var myMap map[int]string
	time.Sleep(time.Millisecond * 150)
	myMap[0] = "golang"

}

// Mutex lock
var count = 0
var mutex sync.Mutex
var m1 = make(map[int]int, 0)

func test1() {
	mutex.Lock()
	count++
	fmt.Println("the count is:", count)
	// time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func test2(num int) {
	mutex.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m1[num] = sum
	fmt.Printf("key=%v value=%v\n", num, sum)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main11() {
	defer func() {
		fmt.Println(count)
	}()
	for i := 1; i < 50; i++ {
		wg.Add(1)
		// go test1()
		go test2(i)
	}

	wg.Wait()
}

var mutexRW sync.RWMutex

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
}

func write() {
	mutexRW.Lock()
	fmt.Println("writing db...")
	time.Sleep(time.Second * 2)
	mutexRW.Unlock()
	wg.Done()
}

func read() {
	mutexRW.RLock()
	fmt.Println("reading db...")
	time.Sleep(time.Second * 2)
	mutexRW.RUnlock()
	wg.Done()
}
