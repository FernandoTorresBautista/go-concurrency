package main

import (
	"fmt"
	"sync"
)

var (
	wg              sync.WaitGroup
	mutex                 = sync.Mutex{}
	widgetInventory int32 = 1000 //Package-level variable will avoid all the pointers
)

// use mutex to solve the problem, exclusive memory with lock and unlock the use of inventory
// expect end = start = 1000000
func main() {
	fmt.Println("Starting inventory count = ", widgetInventory)
	wg.Add(2)
	go makeSales()
	go newPurchases()
	wg.Wait()
	fmt.Println("Ending inventory count = ", widgetInventory)
}

func makeSales() { // 1000000 widgets sold
	for i := 0; i < 300000; i++ {
		mutex.Lock()
		widgetInventory -= 100
		mutex.Unlock()
	}
	wg.Done()
}

func newPurchases() { // 1000000 widgets purchased
	for i := 0; i < 300000; i++ {
		mutex.Lock()
		widgetInventory += 100
		mutex.Unlock()
	}
	wg.Done()
}
