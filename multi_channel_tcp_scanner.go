package main

import (
	"fmt"
	"net"
	"sort"
    "time"
)


//URL_TO_SCAN :="scanme.nmap.org"
var URL_TO_SCAN = "127.0.0.1"
var NB_PORTS = 65535


/*
	The channel will be used to receive work, and the WaitGroup
	to track when a single work item has been completed.
*/
func worker(ports chan int, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("%s:%d", URL_TO_SCAN, p)
		conn, err := net.Dial("tcp", addr)

		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
    start := time.Now()
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= NB_PORTS; i++ {
			ports <- i
		}
	}()

	for i := 0; i < NB_PORTS; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
            fmt.Println("open ", port);
		}
	}
	close(ports)
	close(results)
	sort.Ints(openPorts)

	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
    fmt.Println("time since start: ",time.Since(start))
}
