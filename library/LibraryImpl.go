package library

type LibraryImpl struct {
	storage     BookStorage
	idGenerator IdGenerator
	idByTitle   map[string]BookId
}

func (library *LibraryImpl) GetBook(title string) (*Book, bool) {
	if id, ok := library.idByTitle[title]; ok {
		return library.storage.GetBook(id), true
	} else {
		return nil, false
	}
}

func (library *LibraryImpl) AddBook(book Book) {
	id := library.idGenerator(book, library.storage)
	library.idByTitle[book.Title] = id
	library.storage.AddBook(id, book)
}

func (library *LibraryImpl) SetStorage(storage BookStorage) {
	library.storage = storage
}

func (library *LibraryImpl) SetIdGenerator(idGenerator IdGenerator) {
	library.idGenerator = idGenerator
}
