package main

import (
	"fmt"
	"sync"
)

/*
	The channel will be used to receive work, and the WaitGroup
	to track when a single work item has been completed.
*/
func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100)
	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
