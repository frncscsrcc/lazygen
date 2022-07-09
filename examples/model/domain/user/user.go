// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `User`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `user`

// User: Tish is an user
package user

import (
	"github.com/frncscsrcc/lazygen/examples/model/domain/role"
)

// UserBase is the base struct with the default getters and setters
type UserBase struct {
	id       int64      // Primary key for user
	username string     // User username
	name     string     // User name
	surname  string     // User surname
	password string     // User's password
	phone    []string   // User's phone
	role     *role.Role // User role

	error error
	self  *User
}

// -------------------------------------------

// NewUser returns a pointer to a new NewUser object
func NewUser() *User {
	user := &User{}
	user.phone = make([]string, 0)
	user.self = user
	return user
}

// -------------------------------------------

// Error return the error associate to the object
func (x *UserBase) Error() error {
	return x.error
}

// -------------------------------------------

// Error return the error associate to the object
func (x *UserBase) Validate() error {
	x.error = nil
	return x.error
}

// -------------------------------------------
// Id returns User.id
func (x *UserBase) Id() int64 {
	return x.id
}

// SetId sets User.id
func (x *UserBase) SetId(v int64) *User {
	x.id = v
	return x.self
}

// Username returns User.username
func (x *UserBase) Username() string {
	return x.username
}

// SetUsername sets User.username
func (x *UserBase) SetUsername(v string) *User {
	x.username = v
	return x.self
}

// Name returns User.name
func (x *UserBase) Name() string {
	return x.name
}

// SetName sets User.name
func (x *UserBase) SetName(v string) *User {
	x.name = v
	return x.self
}

// Surname returns User.surname
func (x *UserBase) Surname() string {
	return x.surname
}

// SetSurname sets User.surname
func (x *UserBase) SetSurname(v string) *User {
	x.surname = v
	return x.self
}

// Password returns User.password
func (x *UserBase) Password() string {
	return x.password
}

// SetPassword sets User.password
func (x *UserBase) SetPassword(v string) *User {
	x.password = v
	return x.self
}

// Phone returns []User.phone
func (x *UserBase) Phone() []string {
	return x.phone
}

func (x *UserBase) PhoneIsEmpty() bool {
	return len(x.phone) == 0
}

// PhoneSize returns the size of the array User.phone
func (x *UserBase) PhoneSize() int {
	return len(x.phone)
}

// SetPhone sets User.phone
func (x *UserBase) SetPhone(v []string) *User {
	x.phone = v
	return x.self
}

// AppendPhone append User.phone
func (x *UserBase) AppendPhone(v ...string) *User {
	x.phone = append(x.phone, v...)
	return x.self
}

// Role returns User.role
func (x *UserBase) Role() *role.Role {
	return x.role
}

// SetRole sets User.role
func (x *UserBase) SetRole(v *role.Role) *User {
	x.role = v
	return x.self
}

// -------------------------------------------

// String implements Stringer
func (x *UserBase) String() string {
	jsonString, err := x.ToJSON()
	if err != nil {
		panic(err)
	}
	return "*User" + jsonString
}
