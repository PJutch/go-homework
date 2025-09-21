package library

type BookId int

type BookStorage interface {
	GetBook(id BookId) *Book
	AddBook(id BookId, book Book)
}
