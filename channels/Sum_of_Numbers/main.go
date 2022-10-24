// We will play with channel operations and experiment with their blocking
// nature.
package main

import "log"

// Here, there are two channels, a Boolean one called in, which represents the
// incoming requests, and out, which will be used to send back messages.

func push(from, to int, in <-chan bool, out chan<- int) {
	for i := from; i <= to; i++ {
		<-in //Before sending anything, it waits for a request from the in channel.
		out <- i
	}
}

func main() {
	s1 := 0
	out := make(chan int, 100)
	in := make(chan bool, 100)
	go push(1, 25, in, out)
	go push(26, 50, in, out)
	go push(51, 75, in, out)
	go push(76, 100, in, out)

	// Here, we do not need to sleep for a microsecond because after we receive
	// a number, the next request will go to any active Goroutine.

	for c := 0; c < 100; c++ {
		in <- true //the loop first requests a number
		i := <-out // then waits to receive another number.
		log.Println(i)
		s1 += i
	}
	log.Println(s1)
}
