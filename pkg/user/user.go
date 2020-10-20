package user

type Actor interface {
	Add() error
	Edit() error
	Deactivate() error
}

// User is used to keep a stock of user information.
type User struct {
	Id      string
	Name    Name
	UserId  string
	Role    string // enum for later, possible values = "ADMIN", "LIBRARIAN", "MEMBER"
	Contact Contact
}

type Name struct {
	Salutation string // enum for later, Mr, Mrs, NA
	FirstName  string
	MiddleName string // optional
	LastName   string
}

type Contact struct {
	Email    string // can use validator package to run a regex check
	PhoneNum string // can use validator package to run a regex check
}
