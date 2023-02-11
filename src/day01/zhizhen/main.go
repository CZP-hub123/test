package main

import "fmt"

func changevalue(b *int) {
	*b = 10
}
func main() {
	a := 1
	changevalue(&a)
	fmt.Println("a's value is ", &a)
}
