package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetKind() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleep..\n")

}
func (this *Cat) GetColor() string {
	return this.color

}
func (this *Cat) GetKind() string {
	return "cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog is sleep..\n")

}
func (this *Dog) GetColor() string {
	return this.color

}
func (this *Dog) GetKind() string {
	return "dog"
}

func main() {

	c := Cat{"red"}
	d := Dog{"yellow"}
	c.Sleep()
	fmt.Println(c.GetKind())
	fmt.Println(c.GetColor())
	d.Sleep()
	fmt.Println(d.GetColor())
	fmt.Println(d.GetKind())

	var animal AnimalIF
	animal = &c
	animal.Sleep()
	animal = &d
	animal.Sleep()

}
