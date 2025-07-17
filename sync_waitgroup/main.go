package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() //signal to waitgroup that this goroutine is done
	fmt.Printf("Worker %d starting\n", i)
	//time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
}

func main() {
	//fmt.Println("explore go waitgroup")
	var wg sync.WaitGroup
	//launch 3 goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1) //add 1 to waitgroup
		go worker(i, &wg)
	}
	wg.Wait() //wait for all goroutines to finish
	fmt.Println("all work done")
}
