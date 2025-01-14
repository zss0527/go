package models

import (
	"fmt"
	"time"
)

func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func PrintlnFn(str1, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}
