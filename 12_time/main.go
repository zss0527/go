package main

import (
	"fmt"
	"time"
)

func main() {
	timeObj := time.Now()
	fmt.Println(timeObj)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Println(year, month, day, hour, minute, second)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	//2006:year 01:month 02:day 03:hour(12 clock) 15: hour(24 clock) 04:minute 05:second
	formatTime1 := timeObj.Format("2006-01-02 03:04:05")
	formatTime2 := timeObj.Format("01/02/2006 15:04:05")
	fmt.Println(formatTime1, formatTime2)

	//current timestamp
	unixTime := timeObj.Unix()
	fmt.Println(unixTime)

	//create time from timestamp
	ut := int64(1736586349)
	time1 := time.Unix(ut, 0)
	fmt.Println(time1.Format("01/02/2006 15:04:05"))

	//create time from string
	timeStr := "01/11/2025 17:05:49"
	time2, _ := time.ParseInLocation("01/02/2006 15:04:05", timeStr, time.Local)
	fmt.Println(time2)

	//some funcs for time
	now := time.Now()
	fmt.Println(now.Add(time.Hour))
	//Add Sub Equal Before After

	/*
		timer:
			1. time.NewTicker(time.Second)
			2. time.sleep(time.Second)
	*/
	// ticker := time.NewTicker(time.Second)
	n := 0
	// for v := range ticker.C {
	// 	fmt.Println(v)
	// 	n++
	// 	if n > 5 {
	// 		ticker.Stop()
	// 		return
	// 	}
	// }

	for {
		time.Sleep(time.Second)
		n++
		fmt.Println(time.Now())
		if n > 4 {
			return
		}
	}

}
