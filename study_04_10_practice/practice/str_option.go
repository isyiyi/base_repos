package practice

import "fmt"

func NoneEmpty(s []string) []string{
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i ++
		}
	}
	// 此时，参数和返回值的切片底层共用一个数组，只是len位置不同
	return s[:i]
}

func MapUse() {
	var maps = make(map[string]int)
	maps["wang"] = 90
	maps["li"] = 10
	maps["zjhou"] = 1000
	for name, value := range maps {
		fmt.Println(name, value)
	}
}

func NilCheck() {
	var nums []int
	fmt.Println(nums, nums == nil, len(nums), cap(nums))

	var maps map[string]int
	fmt.Println(maps, maps == nil, len(maps))
}