// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printResult(resultchan chan int, stopgoroutine chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		result := <-resultchan
// 		fmt.Println(result)
// 		stop := <-stopgoroutine
// 		fmt.Println(stop)
// 		if stop {
// 			break
// 		}

// 	}
// }
// func printEven(limit int, oddch chan bool, evench chan bool, resultchan chan int, stopgoroutine chan bool, wg *sync.WaitGroup) {
// 	for i := 2; i <= limit; i += 2 {
// 		<-evench
// 		resultchan <- i
// 		if i == limit {
// 			stopgoroutine <- true
// 		} else {
// 			stopgoroutine <- false
// 			oddch <- true

// 		}
// 	}
// 	wg.Done()
// }
// func printOdd(limit int, oddch chan bool, evench chan bool, resultchan chan int, stopgoroutine chan bool, wg *sync.WaitGroup) {
// 	for i := 1; i <= limit-1; i += 2 {
// 		<-oddch
// 		resultchan <- i
// 		if i == limit-1 {
// 			stopgoroutine <- true
// 		} else {
// 			stopgoroutine <- false
// 			evench <- true
// 		}

// 	}
// 	wg.Done()
// }
// func main() {
// 	var wg sync.WaitGroup
// 	oddch := make(chan bool)
// 	evench := make(chan bool)
// 	resultchan := make(chan int, 10)
// 	stopgoroutine := make(chan bool)
// 	fmt.Println("Enter the value to be printed")
// 	var limit int
// 	fmt.Scanf("%d", &limit)
// 	wg.Add(3)
// 	go printEven(limit, oddch, evench, resultchan, stopgoroutine, &wg)
// 	go printOdd(limit, oddch, evench, resultchan, stopgoroutine, &wg)
// 	go printResult(resultchan, stopgoroutine, &wg)
// 	oddch <- true
// 	wg.Wait()
// }

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
		if stop {
			break
		}
	}
}

func printEven(limit int, oddch chan bool, evench chan bool, resultchan chan int, stopgoroutine chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= limit; i += 2 {
		<-evench
		resultchan <- i
		if i >= limit {
			stopgoroutine <- true
		} else {
			stopgoroutine <- false
			oddch <- true
		}
	}
}

func printOdd(limit int, oddch chan bool, evench chan bool, resultchan chan int, stopgoroutine chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= limit; i += 2 {
		<-oddch
		resultchan <- i
		if i > limit-1 {
			stopgoroutine <- true
		} else {
			stopgoroutine <- false
			evench <- true
		}
	}
}

func main() {
	var wg sync.WaitGroup
	oddch := make(chan bool)
	evench := make(chan bool)
	resultchan := make(chan int, 10)
	stopgoroutine := make(chan bool)

	fmt.Println("Enter the maximum limit to print:")
	var limit int
	fmt.Scanf("%d", &limit)

	wg.Add(3)
	go printEven(limit, oddch, evench, resultchan, stopgoroutine, &wg)
	go printOdd(limit, oddch, evench, resultchan, stopgoroutine, &wg)
	go printResult(resultchan, stopgoroutine, &wg)

	oddch <- true // Start with odd numbers
	wg.Wait()
}
