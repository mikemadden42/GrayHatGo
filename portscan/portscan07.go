package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func scan(ports chan int, host string) {
	for p := range ports {

		conn, err := net.Dial("tcp", host+":"+strconv.Itoa(p))
		if err == nil {
			fmt.Println("local address:", conn.LocalAddr())
			fmt.Println("remote address:", conn.RemoteAddr())
			fmt.Println()

			err = conn.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}
		wg.Done()
	}
}

func main() {
	host := flag.String("host", "localhost", "destination host name")
	debug := flag.Bool("debug", false, "debug mode")
	ports := make(chan int, 100)

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	const maxPorts = 1000

	for i := 0; i < cap(ports); i++ {
		go scan(ports, *host)
	}

	for port := 1; port < maxPorts; port++ {
		if *debug {
			log.Println("DEBUG: testing port", port)
		}
		wg.Add(1)
		ports <- port
	}

	wg.Wait()
	close(ports)
}
