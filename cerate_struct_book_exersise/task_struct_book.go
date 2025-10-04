//Task: Create a struct called Book with Title, Author, Year and print one book.

package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	var book1 Book

	book1.Title = "The Arsène Lupin Collection"
	book1.Author = "Maurice Leblanc"
	book1.Year = 2022

	var book2 Book

	book2.Title = "The Complete Chronicles Of Conan"
	book2.Author = "Robert E Howard"
	book2.Year = 2006

	fmt.Printf("The book is %s %s %d\n", book1.Title, book1.Author, book1.Year)
	fmt.Printf("The book is %s %s %d\n", book2.Title, book2.Author, book2.Year)

}
