package library

type BookStorageImpl struct {
	books     []Book
	indexById map[BookId]int
}

func (storage *BookStorageImpl) GetBook(id BookId) (*Book, bool) {
	if index, ok := storage.indexById[id]; ok {
		return &storage.books[index], true
	} else {
		return nil, false
	}
}

func (storage *BookStorageImpl) AddBook(id BookId, book Book) {
	storage.indexById[id] = len(storage.books)
	storage.books = append(storage.books, book)
}
