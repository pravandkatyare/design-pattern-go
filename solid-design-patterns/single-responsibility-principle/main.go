package main

import "fmt"

type book struct {
	title  string
	author string
}

var books []*book

func (b *book) addBook(bk *book) []*book {
	books = append(books, bk)
	return books
}

func (b *book) display() {
	for _, value := range books {
		fmt.Println(fmt.Sprintf("title: %s, author: %s", value.title, value.author))
	}
}

func main() {
	b := &book{}

	b.addBook(&book{"Design Principles", "XYZ"})
	b.addBook(&book{"Solid Principles", "ABC"})

	b.display()

}
