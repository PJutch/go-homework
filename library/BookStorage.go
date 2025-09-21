package main

type BookId int

type BookStorage interface {
	GetBook(id BookId) (*Book, bool)
	AddBook(id BookId, book Book)
}
