package main

import (
	"fmt"
	"study_04_10_practice/practice"
)

func main() {
	// practice.MainProcess()
	//var nums [3]int = [3]int {2:999}
	//fmt.Println(nums)

	// 关于append内置函数的底层实现逻辑
	//var nums []int
	//for i := 1; i <= 11; i++ {
	//	nums = appendInt(nums, i)
	//	fmt.Println("nums.len=", len(nums), "nums.cap:", cap(nums))
	//	fmt.Println(nums)
	//}


	var nums = make([]int, 2, 4)
	nums[0] = 89
	nums = append(nums, 2, 3, 4, 5, 6, 7)	// 一次性追加多个元素
	var nums2 []int
	nums2 = append(nums2, nums...)					// 直接追加一个切片，但需要拆包，使用copy函数也可以
	fmt.Println(nums)
	fmt.Println(nums2)

	var strs = make([]string, 0, 10)
	strs = append(strs, "wang", "", "", "bing")
	strs = practice.NoneEmpty(strs)
	fmt.Println(strs)

	practice.NilCheck()


	var strs2 string = "中国人民万岁"
	var tmp1 = []byte(strs2)
	fmt.Println(len(tmp1))
	var tmp2 = []rune(strs2)
	fmt.Println(len(tmp2))
}

/* slice的append底层实现:(slice创建时的空间可以直接使用索引方式赋值，使用append的方式都是在原有长度的基础上进行的)
1. 比较source的长度和容量：
	1.1 如果长度小于容量，就在source的底层数组上重新切片，长度较source+1
	1.2 如果长度等于容量，就重新分配一个二倍容量的底层数组，并把source的值赋值过去

*/
func appendInt(source []int, num int) []int {
	var res []int
	if len(source) < cap(source) {
		res = source[:len(source)+1]
	} else {
		if cap(source) == 0 {
			res = make([]int, 1, 4)
		}else{
			res = make([]int, len(source)+1, cap(source)*2)
			for k, v := range source {
				res[k] = v
			}
		}
	}
	res[len(source)] = num
	return res
}
