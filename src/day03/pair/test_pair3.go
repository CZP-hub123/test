package main

import "fmt"

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("read a book")

}
func (this *Book) WriteBook() {
	fmt.Println("write a book")

}

func main() {
	//pair <type :book,value:book's address>
	b := &Book{}

	var r Reader
	//pair <type :book,value:book's address>
	r = b
	r.ReadBook()
	//pair <type :book,value:book's address>
	var w Writer
	w = r.(Writer) //此处断言因为r w 的类型相同
	w.WriteBook()

}
