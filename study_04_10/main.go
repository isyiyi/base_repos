package main

import "fmt"
import "study_04_10/user"

func main() {
	//var f float64 = 212
	//var c float64 = fToC(f)
	//fmt.Println(c)

	// datatype()
	var x int = 1
	// incr(&x)
	fmt.Println(x)

	// new(T)返回的是开辟空间的内存地址，T为指定的数据类型
	var p *int = new(int)
	*p = 98
	fmt.Println(*p)

	typeUse()
	fmt.Println(api.F)
	api.Talk()

	if a, b := returnXAndY(); a == 4 && b == 9 {
		fmt.Println(a, b)
	} else if a == 4 && b == 8 {
		fmt.Println(a, b)
	}

	var height, weight = 82, 180
	if res := bmi(weight, height); res > 40 {
		fmt.Println("重度肥胖")
	} else if res > 35 {
		fmt.Println("中度肥胖")
	} else if res > 30 {
		fmt.Println("轻度肥胖")
	} else if res > 25 {
		fmt.Println("超重")
	} else {
		fmt.Println("正常体重")
	}

	num1, num2 := 8, 8
	fmt.Println((num1 * 1.0) / (num2 * 1.0))

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num1 & num2)
	fmt.Println(num1 | num2)
	fmt.Println(num1 ^ num2)
	fmt.Println(num1 &^ num2)
	num1, num2 = 11, 6
	// &^ 为按位置零，如果第二个操作数的二进制为1，则结果对应的二进制为0，否则对应二进制位的值为第一个操作数的该位的值
	fmt.Println(num1 &^ num2)

	var num3 int = 0x89
	fmt.Printf("%d\n", num3)
	// [1] 表示再次使用第一个操作数
	fmt.Printf("%d\t%[1]d\n", num3)
	fmt.Printf("%d\t%[1]o\t%#[1]x\n", num3)

	var ch1 byte = 'a'
	fmt.Printf("%d\t%[1]c\n", ch1)

	var num4 float64 = 43445.249
	fmt.Printf("%3.4e\n", num4)

	// 复数
	var num5 complex64 = complex(1, 2)
	var num6 complex64 = complex(3, 4)
	fmt.Println(num5 * num6)

	var num7 complex64 = 1 + 2i
	fmt.Println(num7 + num5)

	// len返回的是字符串的字节个数，中文一个字符占3个字节
	var s string = "中华人民共和国"
	fmt.Println(len(s) / 3)

	const num9 int = 90
	const num10 int = 1000
	var num11 = num9 + num10
	fmt.Println(num11)
	num11 = 88
	fmt.Println(num11)

	const (
		KiB = iota
		MiB = 2
		tmp
		tmp_1
		GiB = iota
		TiB
	)
	fmt.Println(KiB, MiB, tmp, tmp_1, GiB, TiB)

	var nums = [...]int{2, 4, 5, 1, 2, 3, 4}
	sliceUse(nums[2:5])
	fmt.Println(nums)
	fmt.Printf("%T\n", nums[:3])

	//var nums3 = make([]int, 9, 10)
	//fmt.Println(len(nums3))
	s2 := "Hello, world, hello golang, hello"
	var tmpS []byte
	for _, v := range s2 {
		tmpS = append(tmpS, byte(v))
	}
	fmt.Printf("%c\t", tmpS)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func datatype() {
	var i, j, k int
	i = 90
	j = 100
	k = i + j
	fmt.Println(k)

	// 指针
	var x int = 100
	// p保存了变量x的内存空间
	var p = &x
	fmt.Println(*p)
	*p = 80
	fmt.Println(x)

}

func incr(p *int) {
	*p++
	fmt.Println(*p)
}

func typeUse() {
	type newType int
	var num1 newType = 90
	var num2 int = 1000
	// print(num1 == num2) 修改名称之后的类型变量与原来类型虽然底层一致，但不再有任何显式的关系
	print(num1 == newType(num2))
}

func returnXAndY() (int, int) {
	return 4, 9
}

func bmi(weight, height int) float64 {
	return float64(weight / height * 2)
}

func sliceUse(nums []int) {
	nums[0] = 999
	nums[1] = 888
	nums[2] = 8989
}
