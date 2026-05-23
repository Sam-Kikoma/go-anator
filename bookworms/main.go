package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var bookPath string
	flag.StringVar(&bookPath,"path","","Enter a valid path eg <path>/<file.json>")
	flag.Parse()

	bookworms, err := loadBookworms(bookPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Here are the books in common:")
	displayBooks(commonBooks)
}

func displayBooks(books []Book){
	for _,book := range books{
		fmt.Println("-", book.Title, "by", book.Author)
	}
}

//stopping point - 3.4.1