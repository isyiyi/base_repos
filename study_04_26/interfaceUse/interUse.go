package interfaceUse

import "fmt"

type ByteCount int

func(c *ByteCount) Write(p []byte)(int, error) {
	fmt.Println("Hello world")
	return len(p), nil
}


func UseInter() {
	var c ByteCount
	fmt.Fprintf(&c, "hello world")
	fmt.Println(c)
}


type USB interface {
	Write(string)(int, error)
	Read(string)(int, error)
	Close() error
}

type Computer struct {
	Brand string
}

type Phone struct {
	Brand string
}

func (p Phone) Write(s string) (int, error) {
	fmt.Println("Phone is writing data...")
	return len(s), nil
}

func (c Computer) Write(s string) (int, error) {
	fmt.Println("computer is writing data...")
	return len(s), nil
}

func (c Computer) Read(s string) (int, error)  {
	fmt.Println(s)
	return len(s), nil
}

func (c Computer) Close() error  {
	fmt.Println("computer is closing...")
	return nil
}

func UseUSB(u USB, s string) {
	u.Write(s)
	u.Read(s)
	u.Close()
}