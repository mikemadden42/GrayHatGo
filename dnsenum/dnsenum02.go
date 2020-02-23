package main

import (
	"fmt"
	"log"
	"net"
)

func resolve(host string) {
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalln(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func main() {
	resolve("microsoft.com")
}
