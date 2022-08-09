package main

import (
	"fmt"
	"strconv"
	"strings"
	"study_04_18/interfaceUse"
)

func test() func(int) {
	var num = 90
	var f func(int) = func(x int) {
		fmt.Println(num)
		num += x
	}
	return f
}

func counter() func() int {
	x := 0
	return func() int {
		x += 1
		return x
	}
}

func geneFix(fix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, fix) {
			return name + fix
		}
		return name
	}
}

//var f func(string)string = geneFix(".jpg")
//res := f("hello")
//fmt.Println(res)
//f = geneFix(".png")
//res = f("hello")
//fmt.Println(res)

func main() {
	//var c interfaceUse.Computer = interfaceUse.Computer{
	//	Name: "联想",
	//}
	//c.Read()
	//c.Write()

	var m Map = make(map[string]int)
	m["adn"] = 89
	fmt.Println(m)

	var u interfaceUse.USB
	var c = interfaceUse.Computer{
		Name: "dell",
	}
	var d = interfaceUse.Computer{
		Name: "apple",
	}
	u = &c
	u.Read()
	u.Write()
	u = &d
	u.Read()
	u.Write()

	//var x interface{} = 90
	//fmt.Println(x)
	//x = "Hello World"
	//fmt.Println(x)
	//x = make(map[string]int)
	// interface{}被称为空接口类型,可以被赋值为任何值,包括函数
	// 但是除了打印,什么都做不了,因为没有空接口没有任何方法
	// 只能使用类型断言之后,才能进行操作
	//var x interface{} = make(map[string]int)
	//x["hello"] = 34

}

type Map map[string]int

func (m Map) String() string{
	var tmp = ""
	for k, v := range m {
		tmp += "[ "
		tmp += k
		tmp += strconv.Itoa(v)
		tmp += "] "
	}
	return tmp
}

