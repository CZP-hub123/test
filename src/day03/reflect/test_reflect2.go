package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "wangjiaer", 90}
	DoFileAndMenthod(user)

}

func DoFileAndMenthod(arg interface{}) {
	fmt.Println("type :", reflect.TypeOf(arg).Name(), "value :", reflect.ValueOf(arg))
	inputType := reflect.TypeOf(arg)
	inputValue := reflect.ValueOf(arg)
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Println("field", field.Name, "value", value)
	}

	//通过type来获取里面的字段
	//1.通过interface获取type，通过type获取numfield，进行便利
	//得到每个field数据类型
	//通过field的interface方法得到对应的value

	//通过type获取里面的方法，
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)

		fmt.Printf("%s:%v\n", m.Name, m.Func)
	}

}
