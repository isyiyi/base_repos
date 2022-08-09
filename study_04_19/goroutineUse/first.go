package goroutineUse

import (
	"fmt"
	"strings"
	"time"
)

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func Fib(x int) int {
	if x < 2 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}

func Gif() {
	for i := 1; i <= 10; i++ {
		tmp := strings.Repeat("■", i)
		fmt.Printf("%s%d%%\r", tmp, i*10)
		time.Sleep(time.Millisecond * 1000)
	}
}


