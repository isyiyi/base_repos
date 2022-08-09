package main

import (
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln("服务器监听8000端口失败...")
	}
	defer listener.Close()
	log.Println("服务器已启动，监听8000端口")

	// 用来存储所有的用户
	var users map[string]net.Conn = make(map[string]net.Conn)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("有连接失败...")
		}

		go handleConn(conn, users)
	}
}

// 启动协程来处理新建立的连接
func handleConn(conn net.Conn, users map[string]net.Conn) {
	defer conn.Close()
	log.Println(conn.RemoteAddr(), "已连接")
	var msg = make([]byte, 1024)

	for {
		n, err := conn.Read(msg)
		var msgTmp = string(msg[:n])
		if err != nil {
			delete(users, msgTmp[12:])
			log.Println(conn.RemoteAddr(), "中断连接")
			break
		}

		if strings.HasPrefix(msgTmp, "用户名：") {
			if _, ok := users[msgTmp]; ok {
				io.WriteString(conn, "该用户已存在")
				continue
			}
			users[msgTmp[12:]] = conn
			io.WriteString(conn, "用户注册完成")
			log.Println("用户注册完成")
			continue
		}
		dealMsg(conn, users, msgTmp)
	}
}

func dealMsg(conn net.Conn, users map[string]net.Conn, msg string) {
	if !strings.Contains(msg, "@") && !strings.Contains(msg, "：") {
		if msg == "list" {
			// 构造所有的用户列表，返回给该用户
			var userList = "所有的用户列表如下：\n"
			for k := range users {
				userList = userList + k + "\n"
			}
			io.WriteString(conn, userList)
		} else {
			io.WriteString(conn, "你的消息服务器已收到")
			return
		}
	} else {
		index1 := strings.IndexAny(msg, "@")
		index2 := strings.IndexAny(msg, "：")
		dest := msg[index1+1 : index2]
		if destConn, ok := users[dest]; !ok {
			io.WriteString(destConn, "目标用户未上线")
		} else {
			mm := msg[index2+3:]
			var srcUser string
			for k, v := range users {
				if v == conn {
					srcUser = k
				}
			}
			io.WriteString(destConn, srcUser+"给您发来消息："+mm)
		}
	}
}
