package library

import (
	"fmt"
)

type Magazine struct {
	BaseItem
	IssueNumber int
}

func (m *Magazine) GetDetails() string {
	return fmt.Sprintf("Magazine - Title: %s, Author: %s, PublishedAt: %s, IssueNumber: %d", m.Title, m.Author, m.PublishedAt, m.IssueNumber)
}
