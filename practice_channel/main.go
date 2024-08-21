package main

import "fmt"

func oddNum(ch <-chan int, done chan<- bool) {
	for num := range ch {
		if num%2 == 0 {
			fmt.Printf("Odd numbers= %d", num)
		}
	}
	done <- true
}

func evenNum(ch <-chan int, done chan<- bool) {
	for num := range ch {
		if num%2 == 0 {
			fmt.Println("Even numbers= %d", num)
		}
	}
	done <- true
}

func main() {

	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	go evenNum(ch, done)
	go oddNum(ch, done)

	<-done
	<-done

}
