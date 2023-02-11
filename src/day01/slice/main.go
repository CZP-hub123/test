package main

import "fmt"

func printArray(myarray [10]int) {

	for i, _ := range myarray {
		fmt.Println(myarray[i])

	}

}

func printArray2(myarray []int) {

	for i, _ := range myarray {
		fmt.Println(myarray[i])

	}

}

func main() {
	var myArray [10]int

	myArray2 := []int{1, 2, 3, 4} //slice  (动态数组)

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])

	}
	for i, i2 := range myArray2 {
		fmt.Println("i=", i, "i2=", i2)
	}

	fmt.Printf("%T", myArray2)

	printArray(myArray)
	printArray2(myArray2)
}
