package main

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type :", reflect.TypeOf(arg), "value :", reflect.ValueOf(arg))

}

func main() {

	var num float64 = 1.2345
	reflectNum(num)
}
