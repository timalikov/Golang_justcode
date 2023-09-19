package library

type LibraryItem interface {
	GetDetails() string
}

type BaseItem struct {
	Title       string
	Author      string
	PublishedAt string
}
