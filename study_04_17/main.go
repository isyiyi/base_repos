package main

import (
	"fmt"
	"study_04_17/knowledge"
)

func main() {
	//var point knowledge.Point = knowledge.Point{X: 23, Y: 90}
	//point.ShowPoint()
	//var nums = make([]int, 2, 4)
	//nums[0] = 9
	//nums[1] = 4
	//point.UpdatePoint(nums...)
	//point.ShowPoint()
	//
	//knowledge.Point{X: 23, Y: 90}.ShowPoint()
	//var pointPtr = &point
	//pointPtr.ShowPoint()
	//(*pointPtr).ShowPoint()

	var point = knowledge.Point{X: 20, Y: 90}
	var point2 = point
	var nums = [...]int {1, 2, 3, 4}
	point2.UpdatePoint(nums[2:4]...)
	point.ShowPoint()
	point2.ShowPoint()


	var x = point.ShowPoint

	point.UpdatePoint(nums[1:3]...)
	x()
	point.ShowPoint()

	knowledge.Point.ShowPoint(point2)

	var m map[string][]int = make(map[string][]int)
	m["hello"] = append(m["hello"], 90)
	fmt.Println(m)


	var point3 knowledge.Point
	point3.ShowPoint()

	var s knowledge.S
	s.Show()

	fmt.Println("*************************")


	var n1 knowledge.Num = 90
	var n2 knowledge.Num = 80

	typeName := "pro"

	// 通过方法表达式语法，将Num类型的方法的接收器放到方法的第一个参数位置
	// 更加灵活的控制方法的接收器和要调用方法（选择器）
	var f func(knowledge.Num, knowledge.Num) knowledge.Num
	switch typeName {
	case "add":
		f = knowledge.Num.Add
	case "sub":
		f = knowledge.Num.Sub
	case "mod":
		f = knowledge.Num.Mod
	case "pro":
		f = knowledge.Num.Pro
	}

	fmt.Println(f(n1, n2))
}
