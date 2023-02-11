package main

import (
	"fmt"
	"time"
)

const (
	Bj = iota
	sh
	sz
)

func main() {
	//var num1 int = 99
	//var num2 float64 = 23.46
	//var str string
	//var b bool = true
	//var myChar byte = 'h'
	//str = fmt.Sprintf("%F", num2)
	//fmt.Printf("str type %T str =%q\n", str, str)
	//
	//str = fmt.Sprintf("%d", num1)
	//fmt.Printf("str type %T str = %q\n", str, str)
	//str = fmt.Sprintf("%t", b)
	//fmt.Printf("str type %T str = %q\n", str, str)
	//
	//str = fmt.Sprintf("%c", myChar)
	//fmt.Printf("str type %T str = %q\n", str, str)

	fmt.Print("hello world")
	time.Sleep(1 * time.Second)

	fmt.Println("beijing  = ", Bj)
	fmt.Println("beijing  = ", sh)
	fmt.Println("beijing  = ", sz)

	var b int = 100
	var c = 100
	fmt.Printf("b type %T c type  %T", b, c)

	//常量

}
