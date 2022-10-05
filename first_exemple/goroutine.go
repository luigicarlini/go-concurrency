package main

import (
	"fmt"
	"time"
)

func msg() {
	//simulating a time consuming process 5 sec
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 1000)
		if i > 3 {
			fmt.Println(i, "second...yawn")
		} else {
			fmt.Println(i, "second...")
		}
	}
	fmt.Println("\n message from function msg(): I have finished!")
}

func main() {
	msg()
	fmt.Println("\n message from function main: I have finished!")
	time.Sleep(time.Millisecond * 2500)
}
