package main

import "fmt"

type Hero struct {
	name  string
	age   int
	level int
} //首字母大写表示共有，小写表示私有

//func (this Hero) getName() string {
//	fmt.Printf("%v\n", this.name)
//	return this.name
//}
////改变不了
//func (this Hero) changeName(newname string) {
//	this.name = newname

// }
func (this *Hero) getName() string {
	fmt.Printf("%v\n", this.name)
	return this.name
}

// 改变不了
func (this *Hero) changeName(newname string) {
	this.name = newname

}

func main() {

	hero := Hero{name: "chaoren", age: 100, level: 1}

	hero.getName()
	hero.changeName("woshi ni die")
	hero.getName()

}
