package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

// 获取标签内容
func findTag(str interface{}) {

	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info", taginfo, "doc", tagdoc)
	}

}

func main() {
	r := resume{"wang", "nv"}
	findTag(&r)
}
