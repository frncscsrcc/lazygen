// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `Role`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `role`

// Role: This is a Role description
package role

import (
	"encoding/json"
)

// DTO to marshal in JSON
type DTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

// -------------------------------------------

func (x *RoleBase) ToDTO() *DTO {
	dto := &DTO{}
	dto.Id = x.id
	dto.Name = x.name
	return dto
}

func (dto *DTO) ToType() *Role {
	role := NewRole()
	role.id = dto.Id
	role.name = dto.Name
	return role
}

// -------------------------------------------

// FromJSON generate a *Role from a valid json (passed as string)

func FromJSON(jsonString string) (*Role, error) {
	dto := &DTO{}
	err := json.Unmarshal([]byte(jsonString), dto)
	if err != nil {
		return NewRole(), err
	}
	role := dto.ToType()
	return role, nil
}

// ToJSON generate a valid JSON (as string) from *Role
func (x *RoleBase) ToJSON() (string, error) {
	dto := x.ToDTO()

	bytes, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
