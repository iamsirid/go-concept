package main

import "fmt"

type Stringer interface {
	String() string
}

type Article struct {
	Title  string
	Author string
}

func (a *Article) String() string {
	return a.Title + " by " + a.Author
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) String() string {
	return b.Title + " by " + b.Author + " with " + fmt.Sprint(b.Pages) + " pages"
}

func main() {

	a := &Article{
		Title:  "Go Interfaces",
		Author: "John Doe",
	}
	Print(a)

	b := &Book{
		Title:  "Go Interfaces2",
		Author: "John Doe",
		Pages:  100,
	}
	Print(b)
}

func Print(s Stringer) {
	fmt.Println(s.String())
}
