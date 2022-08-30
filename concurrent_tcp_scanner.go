package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for j := 1; j <= 65535; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			addr := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(j)
	}
	wg.Wait()
}
