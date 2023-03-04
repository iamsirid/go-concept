package main

import (
	"fmt"
	"sync"
)

func generateNumber(total int, wg *sync.WaitGroup) {
	defer wg.Done()
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumber(wg *sync.WaitGroup) {
	defer wg.Done()
	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go generateNumber(3, &wg)
	wg.Add(1)
	go printNumber(&wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")

	wg.Wait()

}
