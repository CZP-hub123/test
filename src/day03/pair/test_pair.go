package main

import "fmt"

func main() {

	var a string
	//pair<static type :string,value:"abcd">
	a = "abcd"

	//pair<type :string,value:"abcd">
	var alltype interface{}
	alltype = a
	//取万用的type和value
	str, ok := alltype.(string)
	fmt.Println("", str, ok)
	fmt.Printf("%T\n", str)

	//
}
