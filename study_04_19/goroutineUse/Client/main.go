package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatalln(err)
	}
}
