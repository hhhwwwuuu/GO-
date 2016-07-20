/*
	局部变量的作用范围
	由于局部变量，所以子函数f中无法访问变量x
*/

package main

import (
	"fmt"
)

func main() {
	var x string = "hello world"
	fmt.Println(x)
}

func f() {
	fmt.Println(x)
}
