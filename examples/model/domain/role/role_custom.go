// Role: This is a Role description
package role

// Role allows to overwrite the base class
type Role struct {
	RoleBase

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
