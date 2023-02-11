package main

import (
	"fmt"
	"time"
)

// 有无缓冲的通道

func main() {

	c := make(chan int, 3)
	fmt.Println("len(c)=", len(c), "cap(c)", cap(c))
	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在发送数据：", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num=", num)
	}
	fmt.Println("main end")

}

//当channel已经满了，再向里面写数据，就会阻塞
//当channel空的时候，从里面取数据会发生阻塞
