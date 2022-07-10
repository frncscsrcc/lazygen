// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `Address`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `address`

// Address: This entity reppresent an user address
package address

import (
	"encoding/json"
)

// DTO to marshal in JSON
type DTO struct {
	Id      int64  `json:"id"`
	Street  string `json:"street"`
	Number  string `json:"number"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

// -------------------------------------------

func (x *AddressBase) ToDTO() *DTO {
	dto := &DTO{}
	dto.Id = x.id
	dto.Street = x.street
	dto.Number = x.number
	dto.Zip = x.zip
	dto.Country = x.country
	return dto
}

func (dto *DTO) ToType() *Address {
	address := NewAddress()
	address.id = dto.Id
	address.street = dto.Street
	address.number = dto.Number
	address.zip = dto.Zip
	address.country = dto.Country
	return address
}

// -------------------------------------------

// FromJSON generate a *Address from a valid json (passed as string)

func FromJSON(jsonString string) (*Address, error) {
	dto := &DTO{}
	err := json.Unmarshal([]byte(jsonString), dto)
	if err != nil {
		return NewAddress(), err
	}
	address := dto.ToType()
	return address, nil
}

// ToJSON generate a valid JSON (as string) from *Address
func (x *AddressBase) ToJSON() (string, error) {
	dto := x.ToDTO()

	bytes, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
