package main

import "fmt"

func main() {
	//slie := []int{1, 2, 3, 4, 5}
	var slie []int = make([]int, 4)
	var slie2 []int
	fmt.Printf("len = %d ,slice type =%v\n", len(slie2), slie)
	if slie == nil {
		fmt.Println("slice is null")

	} else {
		fmt.Println("slice is not null")
	}

}
