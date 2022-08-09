package interfaceUse

import (
	"bufio"
	"fmt"
	"os"
)

func MainProcess() {
	var f func(a, b int) int
	sc := bufio.NewScanner(os.Stdin)
	switch sc.Scan(); sc.Text() {
	case "+":
		f = add
	case "-":
		f = sub
	case "*":
		f = pro
	default:
		fmt.Println("Hello World")
		return
	}
	var a, b = 90, 80
	res := f(a, b)
	fmt.Println(res)
}

func MainProcess2() {
	fmt.Println()
}



func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func pro(a, b int) int {
	return a * b
}