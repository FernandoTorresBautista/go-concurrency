package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{} // initialize waitgroup

func main() {

	wg.Add(2)            // add two gorutines
	start := time.Now()  // start the time of execution
	go doSomething()     // go rutine 1
	go doSomethingElse() // go rutine 2
	wg.Wait()            // wait for both gorutines

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("\nI've done something")
	wg.Done()
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("I've done something else")
	wg.Done()
}
