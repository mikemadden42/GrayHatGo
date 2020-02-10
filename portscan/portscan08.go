package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
)

func scan(ports, results chan int, host string, debug bool) {
	for p := range ports {
		conn, err := net.Dial("tcp", host+":"+strconv.Itoa(p))
		if err == nil {
			if debug {
				log.Println("local address:", conn.LocalAddr())
				log.Println("remote address:", conn.RemoteAddr())
			}

			err = conn.Close()
			if err != nil {
				log.Fatalln(err)
			}
			results <- p
		} else {
			results <- 0
			continue
		}
	}
}

func main() {
	host := flag.String("host", "localhost", "destination host name")
	debug := flag.Bool("debug", false, "debug mode")
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	const maxPorts = 1000

	for i := 0; i < cap(ports); i++ {
		go scan(ports, results, *host, *debug)
	}

	go func() {
		for port := 1; port <= maxPorts; port++ {
			if *debug {
				log.Println("DEBUG: testing port", port)
			}
			ports <- port
		}
	}()

	for i := 0; i < maxPorts; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
