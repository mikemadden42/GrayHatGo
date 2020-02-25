package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// The list of domains was fetched from https://github.com/bitquark/dnspop.
	subs := subdomains("subdomains.txt")
	domain := flag.String("domain", "microsoft.com", "destination domain name")

	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	resolve(*domain, subs)
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

func resolve(host string, subs []string) {
	for _, sub := range subs {
		ips, err := net.LookupIP(sub + "." + host)
		if err == nil {
			fmt.Println(sub + "." + host)
			for _, ip := range ips {
				fmt.Println(ip)
			}
			fmt.Println()
		}
	}
}
