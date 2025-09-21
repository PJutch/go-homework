package library

type Book struct {
	Title           string
	Author          string
	PublicationYear int
}

func (book Book) FullTitle() string {
	return book.Author + " \"" + book.Title + "\""
}
