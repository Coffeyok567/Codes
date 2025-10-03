// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
)

type Book struct {
	ID         int
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	books map[int]*Book
}

func (l *Library) AddBook(newBook *Book) {
	if l.books == nil {
		l.books = make(map[int]*Book)
	}
	fmt.Printf("%s +\n", newBook.Title)
	l.books[newBook.ID] = newBook
}

func (l *Library) RemoveBook(bookID int) error {
	book, ok := l.books[bookID]
	if ok {
		fmt.Printf("%s -\n", book.Title)
		delete(l.books, bookID)
		return nil
	} else {
		return errors.New("не найдена")
	}
}

func (l *Library) FindByAuthor(author string) []*Book {
	var result []*Book
	for _, book := range l.books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}

func (l *Library) FindByTitle(title string) []*Book {
	var result []*Book
	for _, book := range l.books {
		if book.Title == title {
			result = append(result, book)
		}
	}
	return result
}

func (l *Library) BorrowBook(bookID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("не найдена")
	}
	if book.IsBorrowed {
		return errors.New("уже выдана")
	}
	book.IsBorrowed = true
	fmt.Printf("%s выдана\n", book.Title)
	return nil
}

func (l *Library) ReturnBook(bookID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("не найдена")
	}
	if !book.IsBorrowed {
		return errors.New("уже в библиотеке")
	}
	book.IsBorrowed = false
	fmt.Printf("%s возвращена\n", book.Title)
	return nil
}

func main() {
	library := &Library{}

	book1 := &Book{ID: 1, Title: "Преступление и наказание", Author: "Достоевский"}
	book2 := &Book{ID: 2, Title: "Война и мир", Author: "Толстой"}
	book3 := &Book{ID: 3, Title: "Идиот", Author: "Достоевский"}

	library.AddBook(book1)
	library.AddBook(book2)
	library.AddBook(book3)

	fmt.Println(" Книги Достоевского")
	dostoevskyBooks := library.FindByAuthor("Достоевский")
	for _, book := range dostoevskyBooks {
		fmt.Printf("%s\n", book.Title)
	}

	fmt.Println()
	err := library.BorrowBook(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	err = library.BorrowBook(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	err = library.ReturnBook(1)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println()
	fmt.Println("По имени война и мир")
	warBooks := library.FindByTitle("Война и мир")
	for _, book := range warBooks {
		fmt.Printf("%s, автор: %s\n", book.Title, book.Author)
	}
}
