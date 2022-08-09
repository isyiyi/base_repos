package main

import (
	"fmt"
	"log"
	"os"
	"study_04_11/content"
)

func main() {
	//res := content.StructConJson()
	//content.JsonConvStruct(res)
	//content.GenerateHtml()
	err := content.ErrorUse()
	//fmt.Println(err)

	// 设置输出的前缀格式
	log.SetPrefix("admin:")
	log.SetFlags(0)
	//log.Fatalln(err)

	log.Println(err)
	fmt.Fprintf(os.Stderr, "%s", err)

	var f func() int = content.FuncUse()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

