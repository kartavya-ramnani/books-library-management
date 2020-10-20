package librarian

import "books-library-management/pkg/user"

type Actor interface {
	AddBooks(bookId ...string) error
	RemoveBooks(bookId ...string) error
	ReserveBooks(bookId ...string) error
	BlockUser(userId string) error
	UnblockUser(userId string) error
	UpdateFine(checkoutId string) error
}

type Librarian struct {
	user.User
}

func (l *Librarian) Add() error {
	return nil
}

func (l *Librarian) Edit() error {
	return nil
}

func (l *Librarian) Deactivate() error {
	return nil
}

func (l *Librarian) AddBooks(bookId ...string) error {
	return nil
}

func (l *Librarian) RemoveBooks(bookId ...string) error {
	return nil
}

func (l *Librarian) ReserveBooks(bookId ...string) error {
	return nil
}

func (l *Librarian) BlockUser(userId string) error {
	return nil
}

func (l *Librarian) UnblockUser(userId string) error {
	return nil
}

func (l *Librarian) UpdateFine(checkoutId string) error {
	return nil
}
