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
	host := flag.String("host", "localhost", "destination host name")
	debug := flag.Bool("debug", false, "debug mode")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	const MAX_PORTS = 4096

	for port := 1; port < MAX_PORTS; port++ {
		if *debug {
			log.Println("DEBUG: testing port", port)
		}
		conn, err := net.Dial("tcp", *host+":"+strconv.Itoa(port))
		if err == nil {
			fmt.Println("local address:", conn.LocalAddr())
			fmt.Println("remote address:", conn.RemoteAddr())
			fmt.Println()

			err = conn.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}

	}
}
