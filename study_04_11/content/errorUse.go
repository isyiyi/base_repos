package content

import "fmt"

func ErrorUse() error {
	// 通过fmt.Errorf来格式化一个error，将部分信息注入到原始错误信息中
	return fmt.Errorf("%s", "hello, this is a error")
}


func FuncUse() func() int {
	var x int
	return func() int {
		x += 1
		return x
	}
}