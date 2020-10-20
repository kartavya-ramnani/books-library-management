package admin

import (
	"books-library-management/pkg/user"
	"books-library-management/pkg/user/librarian"
)

type Actor interface {
	AddLibrarian(user user.User) error
	RemoveLibrarian(user user.User) error
}

type Admin struct {
	librarian.Librarian
}

func (a *Admin) Add() error {
	return nil
}

func (a *Admin) Edit() error {
	return nil
}

func (a *Admin) Deactivate() error {
	return nil
}

func (a *Admin) AddLibrarian(user user.User) error {
	return nil
}

func (a *Admin) RemoveLibrarian(user user.User) error {
	return nil
}
