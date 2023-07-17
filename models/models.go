package models

type Book struct {
	id          int
	name        string
	description string
}

type Account struct {
	id   int
	name string
}

func (a *Account) GetID() int {
	return a.id
}

func (a *Account) GetName() string {
	return a.name
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) GetDescriptioN() string {
	return b.description
}

func (a *Account) SetAccountId(id int) {
	a.id = id
}

func (a *Account) SetAccountName(name string) {
	a.name = name
}
