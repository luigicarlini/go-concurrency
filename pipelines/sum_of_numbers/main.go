// We will play with channel operations and experiment with their blocking
// nature.
package main

import (
	"log"
)

// Here, there are two channels, a Boolean one called in, which represents the
// incoming requests, and out, which will be used to send back messages.

func push_1(from, to int) <-chan int {
	out := make(chan int)
	s1 := 0
	go func() {
		for i := from; i <= to; i++ {
			s1 += i
		}
		out <- s1
		close(out)
	}()
	return out
}

func push_2(from, to int, in <-chan int) <-chan int {
	out := make(chan int)
	s2 := <-in
	go func() {
		for i := from; i <= to; i++ {
			s2 += i
		}
		out <- s2
		close(out)
	}()
	return out
}

func push_3(from, to int, in <-chan int) <-chan int {
	out := make(chan int)
	s3 := <-in
	go func() {
		for i := from; i <= to; i++ {
			s3 += i
		}
		out <- s3
		close(out)
	}()

	return out
}

func push_4(from, to int, in <-chan int) <-chan int {
	out := make(chan int)
	s4 := <-in
	go func() {
		for i := from; i <= to; i++ {
			s4 += i

		}
		out <- s4
		close(out)
	}()
	return out
}

func main() {
	// out1 := push_1(1, 25)
	// out2 := push_2(26, 50, out1)
	// out3 := push_3(51, 75, out2)
	// out4 := push_4(76, 100, out3)

	out4 := push_4(76, 100, push_3(51, 75, push_2(26, 50, push_1(1, 25))))

	// Here, we do not need to sleep for a microsecond because after we receive
	// a number, the next request will go to any active Goroutine.

	// for n := range out4 {
	// 	log.Println(n)
	// }
	//log.Println(sum4)
	log.Println(<-out4)
}
