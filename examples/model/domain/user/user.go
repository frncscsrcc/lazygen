// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `User`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `user`

// User: Tish is an user
package user

import (
	"encoding/json"
)

// UserBase is the base struct with the default getters and setters
type UserBase struct {
	id int64 // Primary key for user
	username string // User username
	name string // User name
	surname string // User surname
	password string // User's password
	phone []string // User's phone

	error error
	self *User
}

// userDTO to marshal in JSON
type userDTO struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Password string `json:"password"`
	Phone []string `json:"phone"`
}

// -------------------------------------------

// NewUser returns a pointer to a new NewUser object
func NewUser() *User {
	user := &User{}
	user.self = user
	return user
}

// -------------------------------------------

// Error return the error associate to the object
func (x *UserBase) Error () error {
	return x.error
}

// -------------------------------------------

// Error return the error associate to the object
func (x *UserBase) Validate () error {
	x.error = nil
	return x.error
}

// -------------------------------------------
// Id returns User.id
func (x *UserBase) Id () int64 {
	return x.id
}

// SetId sets User.id
func (x *UserBase) SetId (v int64) *User {
	x.id = v
	return x.self
}
// Username returns User.username
func (x *UserBase) Username () string {
	return x.username
}

// SetUsername sets User.username
func (x *UserBase) SetUsername (v string) *User {
	x.username = v
	return x.self
}
// Name returns User.name
func (x *UserBase) Name () string {
	return x.name
}

// SetName sets User.name
func (x *UserBase) SetName (v string) *User {
	x.name = v
	return x.self
}
// Surname returns User.surname
func (x *UserBase) Surname () string {
	return x.surname
}

// SetSurname sets User.surname
func (x *UserBase) SetSurname (v string) *User {
	x.surname = v
	return x.self
}
// Password returns User.password
func (x *UserBase) Password () string {
	return x.password
}

// SetPassword sets User.password
func (x *UserBase) SetPassword (v string) *User {
	x.password = v
	return x.self
}
// Phone returns []User.phone
func (x *UserBase) Phone () []string {
	return x.phone
}

func (x *UserBase) PhoneIsEmpty () bool {
	return len(x.phone) == 0
}

// PhoneSize returns the size of the array User.phone
func (x *UserBase) PhoneSize () int {
	return len(x.phone)
}

// SetPhone sets User.phone
func (x *UserBase) SetPhone (v []string) *User {
	x.phone = v
	return x.self
}

// AppendPhone append User.phone
func (x *UserBase) AppendPhone (v ...string) *User {
	x.phone = append(x.phone, v...)
	return x.self
}



// -------------------------------------------

// Unmarshal generate a *User from a valid json (bytes)
func Unmarshal (jsonBytes []byte) *User {
	user := &User{}
	dto := &userDTO{}
	err := json.Unmarshal(jsonBytes, dto)
	if err != nil {
		panic("can not unmarshal UserDTO")
	}
	user.id = dto.Id
	user.username = dto.Username
	user.name = dto.Name
	user.surname = dto.Surname
	user.password = dto.Password
	user.phone = dto.Phone
	return user
}

// Mashal generate a valid JSON (bytes) from *User
func (x *UserBase) Marshal () ([]byte, error) {
	dto := userDTO{}
	dto.Id = x.id
	dto.Username = x.username
	dto.Name = x.name
	dto.Surname = x.surname
	dto.Password = x.password
	dto.Phone = x.phone

	return json.Marshal(dto)
}

// -------------------------------------------

// String implements Stringer
func (x *UserBase) String () string {
	jsonBytes, _ := x.Marshal()
	return "*User" + string(jsonBytes)
}
