/*
	go语言中的全局变量的作用范围
*/
package main

import (
	"fmt"
)

var x string = "hello world"

func main() {
	fmt.Println(x + " main")
	f()
}

func f() {
	fmt.Println(x + " f")
}
