package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("人吃饭")
}
func (this *Human) walk() {
	fmt.Println("人z走路")
}

type Superman struct {
	Human //继承父类
	level int
}

func (this *Superman) eat() {
	fmt.Println("超人吃饭") //重写父类方法
}

func (this *Superman) fly() {
	fmt.Println("超人飞翔") //添加子类方法
}
func main() {

	h := Human{name: "zhangsan", sex: "男"}

	h.Eat()
	h.walk()

	s := Superman{Human{"chaoren", "nan"}, 1} //定义子类对象
	s.eat()
	s.walk()
	s.fly()

}
