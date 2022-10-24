package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	// it will first wait for a message and then reply
	msg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello David"
}

func main() {
	ch := make(chan string)
	go greet(ch)
	//Here, the main function is created and a string channel is instantiated.
	//we need to send the first message from the main routine to the second, which is currently waiting.
	ch <- "Hello John"

	//You can see that you need to log twice as you expect two messages to come back.
	//In many cases, you will use a loop to retrieve all the messages
	log.Println(<-ch)
	log.Println(<-ch)
}
