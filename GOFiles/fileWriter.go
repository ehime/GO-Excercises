package main

import (
	"fmt"
	"io/ioutil"
)

// data struct
type Page struct {
	Title string
	Body  []byte									// were []byte is a "byte slice",
}													// REF: http://blog.golang.org/go-slices-usage-and-internals

func (p *Page) save() error {
	filename := p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600) // 0600 permissions
}

func loadPage (title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)			// _ is a blank identifier, essentially assigning error to /dev/null

	return &Page{Title: title, Body: body}
}

func main() {
	p1 := &Page{Title: "Test_page", Body: []byte("Sample page.")}
	p1.save()

	p2 := loadPage("Test_page")
	fmt.Println(string(p2.Body))
}


