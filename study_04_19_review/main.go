package main

import (
	"fmt"
	"time"
)

func main() {
	//ch1 := make(chan int, 5)
	//var ch2 = make(chan int, 5)
	//go counter(ch1)
	//go proPro(ch1, ch2)
	//for v := range ch2 {
	//	fmt.Println(v)

	tick := time.Tick(1 * time.Second)
	stop := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		stop <- struct{}{}
	}()

	i := 10
	for {
		select {
		case <-tick:
			fmt.Println(i)
		case <-stop:
			fmt.Println("stop....")
			return
		}
		i -= 1
	}

}

func counter(ch1 chan<- int) {
	for i := 1; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}

func proPro(ch1 <-chan int, ch2 chan<- int) {
	for v := range ch1 {
		ch2 <- v * v
	}
	close(ch2)
}
