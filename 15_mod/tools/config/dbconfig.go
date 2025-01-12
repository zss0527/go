package config

import "fmt"

func ConfigDB() {
	fmt.Println("db configed...")
}

func init() {
	fmt.Println("inited in config package....")
}
