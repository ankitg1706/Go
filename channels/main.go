package main

import (
	"fmt"
	"sync"
)

func printResult(resultchan chan int, stopgoroutine chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		result := <-resultchan
		fmt.Println(result)
		stop := <-stopgoroutine
		fmt.Println(stop)
		if stop {
			break
		}

	}
}
func printEven(oddch chan bool, evench chan bool, resultchan chan int, stopgoroutine chan bool, wg *sync.WaitGroup) {
	for i := 2; i <= 10; i += 2 {
		<-evench
		resultchan <- i
		if i == 10 {
			stopgoroutine <- true
		} else {
			stopgoroutine <- false
			oddch <- true
			
		}
	}
	wg.Done()
}
func printOdd(oddch chan bool, evench chan bool, resultchan chan int, stopgoroutine chan bool, wg *sync.WaitGroup) {
	for i := 1; i <= 9; i += 2 {
		<-oddch
		resultchan <- i
		// fmt.Println("odd", i)
		stopgoroutine <- false
		evench <- true
		

	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	oddch := make(chan bool)
	evench := make(chan bool)
	resultchan := make(chan int, 10)
	stopgoroutine := make(chan bool)
	wg.Add(3)
	go printEven(oddch, evench, resultchan, stopgoroutine, &wg)
	go printOdd(oddch, evench, resultchan, stopgoroutine, &wg)
	go printResult(resultchan, stopgoroutine, &wg)
	oddch <- true
	wg.Wait()
}
