package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Println("newtask forountine : i= ", i)
		time.Sleep(1 * time.Second)

	}

}

func main() {

	go newTask()

	i := 0
	for {
		i++
		fmt.Println("main de xiancheng", i)
		time.Sleep(1 * time.Second)
	}

}
