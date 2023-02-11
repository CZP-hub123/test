package main

import "fmt"

// interface{}万用类型
func myfunc1(arg interface{}) {
	fmt.Println("myfunc1 is called...", arg)

	//给interface{}提供了“类型断言”机制  向下转型
	//
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string, value=", value)
		fmt.Printf("%T\n", value)
	}
}

type Book struct {
	auth string
}

func main() {
	b := Book{"czp"}
	myfunc1(b)
	myfunc1(100)
	myfunc1("string")
}
