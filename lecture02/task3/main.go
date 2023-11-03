package main

import (
	"fmt"
	"task3/library"
)

func main() {
	book := &library.Book{
		BaseItem: library.BaseItem{
			Title:       "The Go Programming Language",
			Author:      "Alan A. A. Donovan",
			PublishedAt: "2015",
		},
		ISBN:      "9780134190440",
		PageCount: 380,
	}

	magazine := &library.Magazine{
		BaseItem: library.BaseItem{
			Title:       "Tech Monthly",
			Author:      "Tech Media",
			PublishedAt: "September 2023",
		},
		IssueNumber: 42,
	}

	eResource := &library.EResource{
		BaseItem: library.BaseItem{
			Title:       "Learning Go",
			Author:      "Online Academy",
			PublishedAt: "2022",
		},
		URL:  "https://online-academy.com/learning-go",
		Size: 500, // 500 MB
	}

	showDetails(book)
	showDetails(magazine)
	showDetails(eResource)
}

func showDetails(item library.LibraryItem) {
	fmt.Println(item.GetDetails())
}
