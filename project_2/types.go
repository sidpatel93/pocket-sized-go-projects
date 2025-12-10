package main

type byAuthor []Book

func (a byAuthor) Len() int      { return len(a) }
func (a byAuthor) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byAuthor) Less(i, j int) bool {
	if a[i].Author != a[j].Author {
		return a[i].Author < a[j].Author
	}
	return a[i].Title < a[j].Title
}