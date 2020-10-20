package member

import (
	"books-library-management/pkg/user"
	"time"
)

type Actor interface {
	CheckoutBook(bookId string) error
	ReturnBook(bookId string) error
	RenewCheckout(checkoutId string) error
	PayFine(checkoutId string) error
	BuySubscription() error
	RenewSubscription() error
}

type Member struct {
	user.User
	Subs SubscriptionDetail
	// checked out books
	// remaining fines
}

type SubscriptionDetail struct {
	IsSubscribed      bool
	SubscriptionType  string // enum for later
	SubscriptionStart time.Time
	SubscriptionEnd   time.Time
}

func (m *Member) Add() error {
	return nil
}

func (m *Member) Edit() error {
	return nil
}

func (m *Member) Deactivate() error {
	return nil
}

func (m *Member) CheckoutBook(bookId string) error {
	return nil
}

func (m *Member) ReturnBook(bookId string) error {
	return nil
}

func (m *Member) RenewCheckout(checkoutId string) error {
	return nil
}

func (m *Member) PayFine(checkoutId string) error {
	return nil
}

func (m *Member) BuySubscription() error {
	return nil
}

func (m *Member) RenewSubscription() error {
	return nil
}
