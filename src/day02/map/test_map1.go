package main

import "fmt"

func main() {
	var mymap map[string]string

	if mymap == nil {
		fmt.Printf("mymap is null\n")
	}

	mymap = make(map[string]string, 3)
	mymap["one"] = "java"
	mymap["two"] = "go"
	mymap["three"] = "c++"
	fmt.Println(mymap)

	mymap1 := make(map[int]string)

	mymap1[1] = "java"
	mymap1[2] = "C++"

	fmt.Println(mymap1)

	mymap3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "go",
	}
	fmt.Println(mymap3)

}
