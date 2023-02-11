package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	//匿名函数
	go func() {

		defer fmt.Println("A.defer")

		func() {

			defer fmt.Println("B.defer")

			runtime.Goexit() //退出当前goroutine

			fmt.Println("B")

		}()

		fmt.Println("A")

	}()

	for {
		time.Sleep(1 * time.Second)
	}

}
