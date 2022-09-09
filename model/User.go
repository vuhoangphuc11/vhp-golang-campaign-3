package model

import "fmt"

type User struct {
	FullName string
	Age      int
	Email    string
	Phone    string
}

type UserInterface interface {
	GetInfo()
	ChangeInfo()
	InputInfo(u User)
}

func (r User) GetInfo() {
	fmt.Println("User =", r)
}

func (r User) ChangeInfo() {
	r.FullName = "Name Changed"
	fmt.Println("User =", r)
}

func (r *User) InputInfo(u User) *User {
	r.FullName = u.FullName
	r.Age = u.Age
	r.Email = u.Email
	r.Phone = u.Phone
	return r
}
