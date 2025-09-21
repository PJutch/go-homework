package main

import (
	"fmt"
	"math/rand"
)

// Предположим, что в этом задании надо сделать так
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

	if book, ok := library.GetBook("Конец вечности"); ok {
		fmt.Println(book)
	} else {
		panic("this book should be found")
	}

	library.SetIdGenerator(func(title Book, bookStorage BookStorage) BookId {
		return BookId(rand.Int())
	})

	// но ведь смеа генератора id никак не влияет на это
	if book, ok := library.GetBook("Сильмариллион"); ok {
		fmt.Println(book)
	} else {
		panic("this book should be found")
	}

	library.SetStorage(MakeBookStorage())

	library.AddBook(Book{Title: "Преступление и наказание", Author: "Фёдор Михайлович Достоевский", PublicationYear: 1866})
	library.AddBook(Book{Title: "Лев, Колдунья и Платяной шкаф", Author: "Клайвом Стэйплзом Льюисом", PublicationYear: 1950})

	if _, ok := library.GetBook("Сильмариллион"); ok {
		panic("this book should not be found")
	}

	if book, ok := library.GetBook("Преступление и наказание"); ok {
		fmt.Println(book)
	} else {
		panic("this book should be found")
	}
}
