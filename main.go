package main

import (
	"fmt"
	"sync"
	"time"
)

func runIternmitent(end int, seconds int, waiter *sync.WaitGroup) {
	for i := 0; i <= end; i++ {
		var duration int = seconds * int(time.Second)
		time.Sleep(time.Duration(duration))
		fmt.Println(i)
	}
	waiter.Done()
}

func main() {
	fmt.Println("hello again go")
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	waiter.Add(2)
	go runIternmitent(10, 1, waiter)
	go runIternmitent(10, 2, waiter)
	waiter.Wait()
}
