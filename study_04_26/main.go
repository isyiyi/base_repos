package main

import (
	"fmt"
	"net/http"
)

func main() {
	//interfaceUse.UseInter()
	//var c interfaceUse.Computer
	//interfaceUse.UseUSB(c, "this is data")
	//
	//var students = stu{
	//	{id: "S210233", name: "张志华", score: 90.89},
	//	{id: "S210232", name: "王丙超", score: 36},
	//	{id: "S210234", name: "曾柯翔", score: 62},
	//	{id: "S210231", name: "匡俊成", score: 58},
	//	{id: "S210290", name: "孔永康", score: 95},
	//}
	//
	//sort.Sort(students)
	//
	//for k := range students {
	//	fmt.Println(students[k])
	//}
	//
	//x := sort.Reverse(students)
	//sort.Sort(x)
	//for k := range students {
	//	fmt.Println(students[k])
	//}

	//db := database{
	//	"apple":  5999,
	//	"huawei": 4999,
	//	"xiaomi": 2999,
	//}

	//err := errors.New("Hello world")
	//fmt.Println(err)

	//interfaceUse.MainProcess()

}




type Int int
type calc interface {
	Add(int) Int
	Sub(int) Int
}

func (i Int) Add(a int) Int {
	return i + Int(a)
}

func (i Int) Sub(a int) Int{
	return i - Int(a)
}

type Int2 int

func (i Int2) Add(a int) Int {
	return Int(i) + Int(a)
}


type tmp interface {
	Add(int) Int
}




type dollars float32

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		fmt.Println(r.URL.Path)
	case "/getitem":
		item := r.URL.Query().Get("item")
		fmt.Println(item)
		fmt.Println(db[item])
	}
}

func (db database) GetItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	fmt.Println(item)
	fmt.Println(db[item])
}












type Student struct {
	id    string
	name  string
	score float64
}
type stu []Student

func (s stu) Len() int {
	return len(s)
}

func (s stu) Less(i, j int) bool {
	return s[i].score > s[j].score
}

func (s stu) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Student) String() string {
	return fmt.Sprintf("s.id: %s; s.name: %s; s.score: %v", s.id, s.name, s.score)
}
