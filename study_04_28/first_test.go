package main

import (
	"fmt"
	"testing"
)

//// t参数用于报告测试失败和附加的错误信息
//func TestIsPalindrome(t *testing.T) {
//	if !IsPalindrome("Hello world"){
//		t.Errorf("errorInfo: false")
//	}
//}
//
//func TestIsPalindrome2(t *testing.T) {
//	if !IsPalindrome("abba") {
//		t.Errorf("IsPalindrome('abba') = false")
//	}
//}

func TestExam(t *testing.T) {
	var ch = make(chan bool, 1)
	const step = 1e11
	for i := 0; i < 1000000; i ++{
		go func(x int) {
			for j := x * step; j < (x+1) * step; j ++ {
				if j%2 == 1 && j%3 == 2 && j%4 == 1 && j%5 == 4 && j%6 == 5 &&
					j%7 == 4 && j%8 == 1 && j%9 == 2 && j%10 == 9 && j%11 == 0 &&
					j%12 == 5 && j%13 == 10 && j%14 == 11 && j%15 == 14 && j%16 == 9 &&
					j%17 == 0 && j%18 == 11 && j%19 == 18 && j%20 == 9 && j%21 == 11 &&
					j%22 == 11 && j%23 == 15 && j%24 == 17 && j%25 == 9 && j%26 == 23 &&
					j%27 == 20 && j%28 == 25 && j%29 == 16 && j%30 == 29 && j%31 == 27 &&
					j%32 == 25 && j%33 == 11 && j%34 == 17 && j%35 == 4 && j%36 == 29 &&
					j%37 == 22 && j%38 == 37 && j%39 == 23 && j%40 == 9 && j%41 == 1 &&
					j%42 == 11 && j%43 == 11 && j%44 == 33 && j%45 == 29 && j%46 == 15 &&
					j%47 == 5 && j%48 == 41 && j%49 == 46 {
					fmt.Println(j)
				}
			}
		}(i)
	}
	<-ch
}