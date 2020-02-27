package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// The list of domains was fetched from https://github.com/bitquark/dnspop.
	subs := subdomains("subdomains.txt")
	domain := flag.String("domain", "microsoft.com", "destination domain name")
	debug := flag.Bool("debug", false, "debug mode")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	resolveSubDomains(*domain, subs, *debug)
}

func subdomains(file string) []string {
	domains, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer domains.Close()

	scanner := bufio.NewScanner(domains)
	subs := make([]string, 0)
	for scanner.Scan() {
		subs = append(subs, scanner.Text())
	}
	return subs
}

func resolveHost(host string, debug bool) {
	ips, err := net.LookupIP(host)
	if err == nil {
		fmt.Println(host)
		if debug {
			for _, ip := range ips {
				fmt.Println(ip)
			}
			fmt.Println()
		}
	}
	defer wg.Done()
}

func resolveSubDomains(host string, subs []string, debug bool) {
	for _, sub := range subs {
		wg.Add(1)
		go resolveHost(sub+"."+host, debug)
	}
	wg.Wait()
}
