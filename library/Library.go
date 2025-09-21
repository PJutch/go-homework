package library

type IdGenerator func(title Book, bookStorage BookStorage) BookId

type Library interface {
	GetBook(title string) (*Book, bool)
	AddBook(book Book)
	SetStorage(storage BookStorage)
	SetIdGenerator(idGenerator IdGenerator)
}
