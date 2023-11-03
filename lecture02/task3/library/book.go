package library

import (
	"fmt"
)

type Book struct {
	BaseItem
	ISBN      string
	PageCount int
}

func (b *Book) GetDetails() string {
	return fmt.Sprintf("Book - Title: %s, Author: %s, PublishedAt: %s, ISBN: %s, PageCount: %d", b.Title, b.Author, b.PublishedAt, b.ISBN, b.PageCount)
}
