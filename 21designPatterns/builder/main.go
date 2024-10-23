package main

import "fmt"

type User struct {
	fName string
	lName string
	add   string
	phone string
}

type UserBuilder struct {
	user *User
}

func NewUserBuilder(fName, lName string) *UserBuilder {
	return &UserBuilder{
		&User{
			lName: lName,
			fName: fName,
		},
	}
}

func (ub *UserBuilder) StaysAt(add string) *UserBuilder {
	ub.user.add = add
	return ub
}
func (ub *UserBuilder) CallAt(phone string) *UserBuilder {
	ub.user.phone = phone
	return ub
}

func (ub *UserBuilder) Build() *User {
	return ub.user
}

func main() {
	user := NewUserBuilder("Foo", "Bar").
		StaysAt("Banglore").
		CallAt("9999999").
		Build()

	fmt.Println("user: ", user)

}
