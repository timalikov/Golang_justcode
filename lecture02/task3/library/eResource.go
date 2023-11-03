package library

import (
	"fmt"
)

type EResource struct {
	BaseItem
	URL  string
	Size int // Size in MB
}

func (e *EResource) GetDetails() string {
	return fmt.Sprintf("E-Resource - Title: %s, Author: %s, PublishedAt: %s, URL: %s, Size: %dMB", e.Title, e.Author, e.PublishedAt, e.URL, e.Size)
}
