package interfaceUse

import "fmt"

type USB interface {
	Read()
	Write()
	Closer()
}

type Computer struct {
	Name string
}

func (c *Computer) Read() {
	fmt.Println(c.Name)
	fmt.Println("read...")
}

func (c *Computer) Write() {
	fmt.Println(c.Name)
	fmt.Println("write...")
}

func (c *Computer) Closer() {
	fmt.Println("close")
}