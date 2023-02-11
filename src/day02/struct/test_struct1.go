package main

import "fmt"

type Bool struct {
	title string
	auth  string
}

// 指针
func changgebooktitle(book *Bool) {
	book.title = "天地大同"
}

func main() {

	book1 := Bool{title: "Goland", auth: "chenzhipeng"}

	fmt.Printf("%v\n", book1)
	changgebooktitle(&book1)
	fmt.Printf("%v\n", book1)

	hero := Hero{name: "die", age: 10, level: 1}

	fmt.Printf("%v\n", hero)

}
