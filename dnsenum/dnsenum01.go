package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ips, err := net.LookupIP("microsoft.com")
	if err != nil {
		log.Fatalln(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
