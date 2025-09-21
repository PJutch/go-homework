package main

// Scenario to test the library bcs idk how to do unit-tests here
func main() {
	books := []Book{
		{Title: "Война и Мир", Author: "Лев Николаевич Толстой", PublicationYear: 1865},
		{Title: "Конец вечности", Author: "Айзек Азимов", PublicationYear: 1955},
		{Title: "Сильмариллион", Author: "Джон Рональд Руэл Толкин", PublicationYear: 1977}}

	lastId := 0
	var library Library = MakeLibrary(func(title Book, bookStorage BookStorage) BookId {
		lastId += 1
		return BookId(lastId)
	})

	for _, book := range books {
		library.AddBook(book)
	}
}
