package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("local address:", conn.LocalAddr())
	fmt.Println("remote address:", conn.RemoteAddr())

	err = conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
