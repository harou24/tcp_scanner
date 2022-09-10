package main

import (
	"fmt"
	"net"
    "time"
)

var URL_TO_SCAN = "127.0.0.1"
var NB_PORTS = 65535

func main() {
    start := time.Now()
	for i := 1; i <= NB_PORTS; i++ {
		addr := fmt.Sprintf("%v:%d",URL_TO_SCAN, i)

		// the Dial function connects to a server.
		conn, err := net.Dial("tcp", addr)

		if err != nil {
			// port is closed or filtered.
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
    fmt.Println("Time since start: ", time.Since(start))
}
