package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ips, err := net.LookupIP("www.google.com")
	if err != nil {
		log.Fatalln(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
