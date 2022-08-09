package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex
var count int

var name string = "golang"

// 一个文件中可以有多个init函数，每个init函数都会在文件开始时自动被调用，且顺序按照声明顺序
func init() {
	fmt.Println("Hello")
}

func init() {
	fmt.Println("Golang")
}

type A struct {
	X int
	Y int
}

type C struct {
	X, Y int
}

type B struct {
	A
	C
	X int
}

func main() {
	//var url string = "www.bing.com/"
	//getContent(url)

	//http.HandleFunc("/", handler)
	//http.HandleFunc("/count", counter)
	//
	//log.Fatal(http.ListenAndServe("localhost:8080", nil))

	// 设置 命令行参数名字和默认值、描述信息
	//// 输入-h or -help来显式描述信息
	//var n = flag.Bool("n", false, "n is invalid")
	//var seq = flag.String("seq", "_", "seq is invalid")

	// 参数格式化，将os.Args()[1:]参数的值更新为用户输入的命令行参数，默认是上面写的默认值
	// 除了上述两个命令行参数外，还可以有其他的
	//flag.Parse()
	//fmt.Println(strings.Join(flag.Args(), *seq))
	//if *n {
	//	fmt.Println("Hello")
	//}
	//
	//fmt.Println(gcd(20, 8))
	//
	//fmt.Println(ss.Add(23, 23))

	if len(os.Args) > 4 || len(os.Args) < 2 {
		fmt.Println("error")
	} else {
		tmpA, _ := strconv.Atoi(os.Args[1])
		tmpB, _ := strconv.Atoi(os.Args[2])
		fmt.Printf("%s + %s = %d\n", os.Args[1], os.Args[2], tmpA+tmpB)
	}

	x := "hello"
	for i := 0; i < len(x); i++ {
		x := x[i] + 'A' - 'a'
		fmt.Printf("%c", x)
	}

	num1 := 90
	for i := 0; i < 10; i++ {
		num1 := 10 + num1
		fmt.Println(num1)
	}
	var name string = "中华人民共和国"
	nameByte := []byte(name)
	nameRune := []rune(name)
	fmt.Printf("%c\n%c\n", nameByte, nameRune)
	var address string = "New York"
	addrByte := []uint8(address)
	fmt.Printf("%c\n", addrByte)

	var num4 int = 90
	var num2 int = 11
	var num3 = num4 * num2
	var num5 = 43 * 9.8
	fmt.Println(num3)
	fmt.Printf("%T;%[1]v\n", num5)

	for i := 0; i < len([]rune(name)); i++ {
		fmt.Println(name[i])
	}

	const (
		_ = 1 << (10 * iota)
		a
		b
	)
	fmt.Println(1 << 10)
	fmt.Println(b)

	var bb B
	bb.X = 90
	bb.A.X = 80
	bb.C.X = 100
	bb = B{A: A{X: 1, Y: 34}, C: C{X: 234, Y: 90}, X: 23}
	fmt.Printf("%v\n", bb)
	bb = B{A{12, 23}, C{23, 34}, 90}
	fmt.Printf("%#v\n", bb)

	var aaa int = 90
	var s = getB(aaa)
	s.A.X = 90
	fmt.Println(s)

	// 进栈出栈，所以顺序颠倒
	// defer语句会在所在的函数执行完毕之后执行，并且执行顺序与声明顺序相反
	defer fmt.Println(2)


	fmt.Println("Hello golang")
	defer fmt.Println(3)
	defer fmt.Println(4)

	fmt.Scan(&num2)


	defer func() {
		p := recover()
		fmt.Println("Hello:", p)
	}()
}

func getB(id int) B {
	var bb B
	bb.X = id
	return bb
}


func getContent(url string) {
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("网页有点问题")
		return
	}
	var content []byte
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err == io.EOF {
		fmt.Println("读到文件结尾")
	}
	fmt.Println(resp.Status)
	resp.Body.Close()
	fmt.Println(string(content))

}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path:%s %s %s\n", r.URL.Path, r.Method, r.Proto)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count:%d\n", count)
	mu.Unlock()
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
