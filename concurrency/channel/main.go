package main

import (
	"fmt"
	"sync"
)

func generateNumber(total int, ch chan<- int) {
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Sending number %d to channel\n", idx)
		ch <- idx
	}
}

func printNumber(idx int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("%d: Printing number %d from channel\n", idx, num)
	}
}

func main() {

	var wg sync.WaitGroup
	numberChan := make(chan int)

	for idx := 1; idx <= 3; idx++ {
		wg.Add(1)
		go printNumber(idx, numberChan, &wg)
	}

	generateNumber(5, numberChan)

	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")

	wg.Wait()

}
