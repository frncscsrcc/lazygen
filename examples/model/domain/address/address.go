// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `Address`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `address`

// Address: This entity reppresent an user address
package address

// AddressBase is the base struct with the default getters and setters
type AddressBase struct {
	id int64
	street string
	number string
	zip string
	country string

	error error
	self *Address
}

// -------------------------------------------

// NewAddress returns a pointer to a new NewAddress object
func NewAddress() *Address {
	address := &Address{}
	address.self = address
	return address
}

// -------------------------------------------

// Error return the error associate to the object
func (x *AddressBase) Error () error {
	return x.error
}

// -------------------------------------------

// Error return the error associate to the object
func (x *AddressBase) Validate () error {
	x.error = nil
	return x.error
}

// -------------------------------------------
// Id returns Address.id
func (x *AddressBase) Id () int64 {
	return x.id
}

// SetId sets Address.id
func (x *AddressBase) SetId (v int64) *Address {
	x.id = v
	return x.self
}
// Street returns Address.street
func (x *AddressBase) Street () string {
	return x.street
}

// SetStreet sets Address.street
func (x *AddressBase) SetStreet (v string) *Address {
	x.street = v
	return x.self
}
// Number returns Address.number
func (x *AddressBase) Number () string {
	return x.number
}

// SetNumber sets Address.number
func (x *AddressBase) SetNumber (v string) *Address {
	x.number = v
	return x.self
}
// Zip returns Address.zip
func (x *AddressBase) Zip () string {
	return x.zip
}

// SetZip sets Address.zip
func (x *AddressBase) SetZip (v string) *Address {
	x.zip = v
	return x.self
}
// Country returns Address.country
func (x *AddressBase) Country () string {
	return x.country
}

// SetCountry sets Address.country
func (x *AddressBase) SetCountry (v string) *Address {
	x.country = v
	return x.self
}

// -------------------------------------------

// String implements Stringer
func (x *AddressBase) String () string {
	jsonString, err := x.ToJSON()
	if err != nil {
		panic(err)
	}
	return "*Address" + jsonString
}
