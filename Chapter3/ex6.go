/*
	constants 的应用
	const 是一个被预设的变量，在之后的执行中，该变量将无法被改变。
*/

package main

import (
	"fmt"
)

func main() {
	const x string = "hello world"

	fmt.Println(x)
	x = "ex6"
}
