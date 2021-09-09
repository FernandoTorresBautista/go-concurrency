package main

import (
	"fmt"
)

func main() {

	c := make(chan string, 3) // channel doesn't block until full ("buffered" channel)
	c <- "Hello "             // 1
	c <- "Earth "             // 2
	c <- "from Mars"          // 3
	// 4 here there is a deadlock 'cause are 3 items in the channel
	// c <- "from Venus"

	msg := <-c
	fmt.Print(msg)

	msg = <-c // Notice we used = NOT := because msg is already declared
	fmt.Print(msg)

	msg = <-c // Notice we used = NOT := because msg is already declared
	fmt.Println(msg)

}
