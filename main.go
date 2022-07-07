package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	processed := make(chan string)
	subdomains := make(chan string)
	domains := make(map[string]struct{})
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			for subdomain := range subdomains {
				if (strings.Contains(subdomain, "Discovered open port")) {
					words := strings.Fields(subdomain)
					record := words[5] + ":" + words[3]
					clean := strings.Split(record, "/")
					record = clean[0]
					processed <- record
				}

			}

			wg.Done()
		}()
	}

	go func() {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			subdomains <- sc.Text()
		}
		close(subdomains)
	}()
	go func() {
		wg.Wait()
		close(processed)
	}()
	for domain := range processed {
		_, ok := domains[domain]
		if !ok {
			domains[domain] = struct{}{}
			fmt.Println(domain)
		}
	}

}
