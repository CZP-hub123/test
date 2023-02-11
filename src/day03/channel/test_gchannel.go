package main

import (
	"fmt"
)

// 关闭channel
func main() {
	c := make(chan int)
	fmt.Println("len(c)=", len(c), "cap(c)", cap(c))
	go func() {
		defer fmt.Println("子go程结束")
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("子go程正在发送数据：", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
		//	关闭channel
		close(c)
	}()

	//for {
	//	//未关闭会出现死锁
	//	if value, ok := <-c; ok {
	//		fmt.Println(value)
	//	} else {
	//		break
	//	}
	//
	//}
	//可以用range去读取通道中的数据
	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("main end")

}
