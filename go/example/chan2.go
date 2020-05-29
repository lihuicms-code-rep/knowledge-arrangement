package main

import (
	"fmt"
	"time"
)

//通道同步

func worker(done chan bool) {
	fmt.Println("working......")
	time.Sleep(2 * time.Second)
	fmt.Println("done......")
	done <- true
}

func waiter(done chan bool) {
	fmt.Println("begin wait.....")
	<-done
	fmt.Println("ok, wait")
}
func main() {
	done := make(chan bool)
	go worker(done)
	go waiter(done)
	select {

	}
}
