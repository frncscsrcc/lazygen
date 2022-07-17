// User: This is an user
package user

// User allows to overwrite the base class
type User struct {
	UserBase

	// add public or private fields here
}

/*
Overwrite functions defined in ./user.go
Eg:

// SetUsername sets User.username
func (x *User) SetFieldX (s string) *User {
	x.fieldX = s + " something else"
	return x.self
}
*/
