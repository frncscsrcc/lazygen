// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `User`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `user`

// User: Tish is an user
package user

import (
	"encoding/json"
	"github.com/frncscsrcc/lazygen/examples/model/domain/address"
	"github.com/frncscsrcc/lazygen/examples/model/domain/role"
)

// DTO to marshal in JSON
type DTO struct {
	Id        int64          `json:"id"`
	Username  string         `json:"username"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	Password  string         `json:"password"`
	Phone     []string       `json:"phone"`
	Role      *role.DTO      `json:"role"`
	Addresses []*address.DTO `json:"addresses"`
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
	if x.role != nil {
		dto.Role = x.role.ToDTO()
	}
	dto.Addresses = make([]*address.DTO, 0)
	for _, address := range x.addresses {
		if address != nil {
			dto.Addresses = append(dto.Addresses, address.ToDTO())
		}
	}
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
	if dto.Role != nil {
		user.role = dto.Role.ToType()
	}
	user.addresses = make([]*address.Address, 0)
	for _, address := range dto.Addresses {
		if address != nil {
			user.addresses = append(user.addresses, address.ToType())
		}
	}
	return user
}

// -------------------------------------------

// FromJSON generate a *User from a valid json (passed as string)

func FromJSON(jsonString string) (*User, error) {
	dto := &DTO{}
	err := json.Unmarshal([]byte(jsonString), dto)
	if err != nil {
		return NewUser(), err
	}
	user := dto.ToType()
	return user, nil
}

// ToJSON generate a valid JSON (as string) from *User
func (x *UserBase) ToJSON() (string, error) {
	dto := x.ToDTO()

	bytes, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
