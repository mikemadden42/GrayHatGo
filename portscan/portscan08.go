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

func scan(ports, results chan int, host string) {
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
			results <- p
		} else {
			results <- 0

		}
	}
}

func main() {
	host := flag.String("host", "localhost", "destination host name")
	debug := flag.Bool("debug", false, "debug mode")
	ports := make(chan int, 10)
	results := make(chan int)
	var openports []int

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	const maxPorts = 100

	for i := 0; i < cap(ports); i++ {
		go scan(ports, results, *host)
	}

	go func() {
		for port := 1; port < maxPorts; port++ {
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
