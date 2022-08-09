package knowledge

import "fmt"

type Point struct {
	X int
	Y int
}


func (p Point) ShowPoint() {
	fmt.Println(p.X, p.Y)
}

func (p *Point) UpdatePoint(nums ...int) {
	if len(nums) == 0 {
		fmt.Println("无更新变量")
		return
	}
	(*p).X = nums[0]
	(*p).Y = nums[1]
}

func (p Point) UpdatePoint2(nums ...int) {
	if len(nums) == 0 {
		fmt.Println("无更新变量")
		return
	}
	p.X = nums[0]
	p.Y = nums[1]
}

type Pio *Point

type S string

func (ss S) Show() {
	fmt.Println(ss)
}

type Values map[string][]int
