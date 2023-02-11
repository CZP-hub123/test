package main

import "fmt"

//	func fool(a string, b int) int {
//		fmt.Println("a = ", a)
//		fmt.Println("b=", b)
//
//		c := 100
//		return c
//	}
func fool2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b=", b)

	c := 100
	d := 90
	return c, d
}

//函数多返回值

func main() {

	c, d := fool2("werer", 100)
	fmt.Println("c= ", c, "d =", d)

}
