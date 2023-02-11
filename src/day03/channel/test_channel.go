package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine run ...")

		c <- 666
		fmt.Println("before defer")

	}()
	ok := <-c
	fmt.Println(ok)

}
