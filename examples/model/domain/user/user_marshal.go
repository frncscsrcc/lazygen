// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `User`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `user`

// User: Tish is an user
package user

import (
	"encoding/json"
    "errors"
	"github.com/frncscsrcc/lazygen/examples/model/domain/role"
)

// DTO to marshal in JSON
type DTO struct {
    Id int64 `json:"id"`
    Username string `json:"username"`
    Name string `json:"name"`
    Surname string `json:"surname"`
    Password string `json:"password"`
    Phone []string `json:"phone"`
    Role *role.DTO `json:"role"`
}

// -------------------------------------------

func (x *UserBase) ToDTO() *DTO {
	dto := &DTO{}
	dto.Id = x.id
	dto.Username = x.username
	dto.Name = x.name
	dto.Surname = x.surname
	dto.Password = x.password
	dto.Phone = x.phone
	dto.Role = x.role.ToDTO()
	return dto
}

func (dto *DTO) ToType() *User {
    user := NewUser()
	user.id = dto.Id
	user.username = dto.Username
	user.name = dto.Name
	user.surname = dto.Surname
	user.password = dto.Password
	user.phone = dto.Phone
	user.role = dto.Role.ToType()
	return user
}


// -------------------------------------------

// FromJSON generate a *User from a valid json (bytes)

func FromJSON (jsonBytes []byte) (*User, error) {
	user := &User{}
	dto := &DTO{}
	err := json.Unmarshal(jsonBytes, dto)
	if err != nil {
        return NewUser(), errors.New("can not unmarshal UserDTO")
	}
	user.id = dto.Id
	user.username = dto.Username
	user.name = dto.Name
	user.surname = dto.Surname
	user.password = dto.Password
	user.phone = dto.Phone
	user.role = dto.Role.ToType()
	return user, nil
}


// ToJSON generate a valid JSON (bytes) from *User
func (x *UserBase) ToJSON () (string, error) {
	dto := DTO{}
	dto.Id = x.id
	dto.Username = x.username
	dto.Name = x.name
	dto.Surname = x.surname
	dto.Password = x.password
	dto.Phone = x.phone
	dto.Role = x.role.ToDTO()

	bytes, err := json.Marshal(dto)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

