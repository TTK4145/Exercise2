// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

func numbeServer(addNumber <-chan int, control <-chan int, number chan<- int) {
	var i = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
		case update := <-addNumber:
			i += update

		case signal := <-control:
			if signal == GetNumber {
				number <- i
			}
			if signal == Exit {
				return
			}

			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.
		}
	}
}

func incrementing(addNumber chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000; j++ {
		addNumber <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementing(addNumber chan<- int, finished chan<- bool) {
	for j := 0; j < 1000000-1; j++ {
		addNumber <- -1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
	finished := make(chan bool, 2)
	number := make(chan int)
	addNumber := make(chan int)
	control := make(chan int)

	// TODO: Spawn the required goroutines
	go numbeServer(addNumber, control, number)
	go incrementing(addNumber, finished)
	go decrementing(addNumber, finished)

	<-finished
	<-finished

	// TODO: block on finished from both "worker" goroutines

	control <- GetNumber
	Println("The magic number is:", <-number)
	control <- Exit
}
