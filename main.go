package main

import (
	"custom-wait/waitsync"
	"fmt"
	"time"
)

func main() {
	var wg = waitsync.NewWaitGroup() // create an instance of the customised WaitGroup

	// To create on the fly 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(i))
			fmt.Printf("Goroutine %d just finished\n", i)

		}(i)
	}

	wg.Wait() // Add Wait to block till the counter is zero

	fmt.Println("All goroutine successfully finished!")

}
