package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	host := flag.String("host", "scanme.nmap.org", "destination host name")
	port := flag.Int("port", 80, "destination port")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", *host+":"+strconv.Itoa(*port))
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
