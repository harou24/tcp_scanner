package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 65535; i++ {
		addr := fmt.Sprintf("127.0.0.1:%d", i)

		// the Dial function connects to a server.
		conn, err := net.Dial("tcp", addr)

		if err != nil {
			// port is closed or filtered.
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
