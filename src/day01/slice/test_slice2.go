package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("len=%d, cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 2)
	fmt.Printf("len=%d, cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 3)
	fmt.Printf("len=%d, cap=%d,slice=%v\n", len(numbers), cap(numbers), numbers)

	s := numbers[:5]

	s2 := make([]int, 5)
	copy(s, s2)

	fmt.Printf("len=%d, cap=%d,slice=%v\n", len(s), cap(s), s)
	s[0] = 10
	fmt.Println(s[0])
	fmt.Println(s2)
}
