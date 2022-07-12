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

// NewAddress build a DTO empty structure
func NewDTO() *DTO {
	dto := &DTO{}
	return dto
}

// -------------------------------------------

func (x *AddressBase) ToDTO() *DTO {
	dto := NewDTO()
	dto.Id = x.Id()
	dto.Street = x.Street()
	dto.Number = x.Number()
	dto.Zip = x.Zip()
	dto.Country = x.Country()
	return dto
}

func (dto *DTO) ToType() *Address {
	address := NewAddress()
	address.SetId(dto.Id)
	address.SetStreet(dto.Street)
	address.SetNumber(dto.Number)
	address.SetZip(dto.Zip)
	address.SetCountry(dto.Country)
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
