package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	if err != nil {
		log.Fatalln("连接8000端口失败")
	}
	input := bufio.NewScanner(os.Stdin)
	var name string
	var msg string
	var res = make([]byte, 1024)
	for {
		fmt.Println("请输入你的用户名：")
		if input.Scan() {
			name = input.Text()
		}
		_, err = io.WriteString(conn, "用户名：" + name)
		if err != nil {
			log.Fatalln("客户端发送消息失败")
		}
		n, err := conn.Read(res)
		if err != nil {
			log.Fatalln("接收消息失败")
		}
		if string(res[:n]) == "用户注册完成" {
			break
		}
	}


	var msgs = make(chan string, 10)

	go receive(conn, msgs)
	go printMsg(msgs)

	for {
		fmt.Printf("%s，您好，请输入消息：\n", name)
		if input.Scan() {
			msg = input.Text()
		}
		_, err = io.WriteString(conn, msg)
		if err != nil {
			log.Fatalln("客户端发送消息失败")
		}
	}
}


func receive(conn net.Conn,  msgs chan <- string) {
	var buf = make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatalln("服务器已关闭")
		}
		msgs <- string(buf[:n])
	}
}

func printMsg(msgs <-chan string) {
	for {
		var msg = <- msgs
		fmt.Println(msg)
	}
}