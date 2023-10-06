package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go getOrder(i, &wg)
	}

	wg.Wait()
	fmt.Println("Finished")
	fmt.Println(time.Since(start))
}

func getOrder(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	//time.Sleep(1 * time.Second)
	fmt.Printf("The API call, number %d, was executed\n", id)
}
