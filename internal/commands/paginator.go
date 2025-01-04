package commands

// Paginator handles pagination logic for the location areas
type Paginator struct {
	pageSize    int
	currentPage int
}

// Package level paginator instance
var PaginatorInstance *Paginator

// NewPaginator creates a new Paginator starting at page 1
func NewPaginator(pageSize int) *Paginator {
	return &Paginator{
		pageSize:    pageSize,
		currentPage: 1,
	}
}

// GetCurrentPageItems returns the start and end item numbers for the current page
func (p *Paginator) GetCurrentPageItems() (start, end int) {
	start = ((p.currentPage - 1) * p.pageSize) + 1
	end = start + p.pageSize - 1
	return start, end
}

// NextPage moves to the next page and returns the new page's start and end items
func (p *Paginator) NextPage() (start, end int) {
	start, end = p.GetCurrentPageItems()
	p.currentPage++
	return start, end
}

// PreviousPage moves to the previous page and returns the new page's start and end items
// Returns false if already on the first page
func (p *Paginator) PreviousPage() (start, end int, ok bool) {
	if p.currentPage <= 1 {
		p.currentPage = 1
		start, end = p.GetCurrentPageItems()
		return start, end, false
	}
	p.currentPage--
	start, end = p.GetCurrentPageItems()
	return start, end, true
}

// GetCurrentPage returns the current page number
func (p *Paginator) GetCurrentPage() int {
	return p.currentPage
}
