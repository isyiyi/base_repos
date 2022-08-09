package main

import (
	"fmt"
)

//func IsPalindrome(s string) bool {
//	for i := range s {
//		if s[i] != s[len(s)-1-i] {
//			return false
//		}
//	}
//	return true
//}
//

func main() {
	//var a = "Hello world"
	//t := reflect.TypeOf(a)
	//fmt.Println(t.String())
	//
	//var b = 9
	//x := reflect.ValueOf(&b)
	//d := x.Elem()
	//fmt.Println(d.CanAddr())
	//
	//if d.CanAddr() {
	//	y := d.Addr().Interface().(*int)
	//	*y = 10
	//}
	//fmt.Println(b)
	//var c = make([]int, 10)
	//c[0] = 90
	//c[1] = 100
	//c[2] = 80
	//e := reflect.ValueOf(c)
	//w := e.Index(0)
	//fmt.Println(w.CanAddr())
	//var balance money = 1000
	//go (&balance).Save(1000)
	//go (&balance).Get(800)
	//fmt.Println(balance)

	//var count = 0
	//var mu sync.Mutex
	////mu.Lock()
	////mu.Unlock()
	//
	////var lock = make(chan int, 1)
	//for i := 0; i < 10000; i++ {
	//	go func() {
	//		mu.Lock()
	//		defer mu.Unlock()
	//		count += 1
	//	}()
	//}
	//time.Sleep(1 * time.Second)
	//fmt.Println(count)

	var bal int
	var ch = make(chan int)
	var ch2 = make(chan int)

	go func() {
		for {
			select {
			case bal = <-ch:
				fmt.Println(bal)
			case ch2 <- bal:
			}
		}
	}()
	ch <- 90
	fmt.Println(<-ch2)



}

type money int

func (m *money) Save(moneys money) {
	fmt.Println("Save, balance = ", *m)
	*m += moneys
}

func (m *money) Get(moneys money) {
	fmt.Println("Get, balance = ", *m)
	*m -= moneys
}
