// THIS CODE IS AUTO-GENERATED DO NOT MANUALLY EDIT THIS FILE!!!!
// IF YOU NEED TO OVERWRITE OR ADD METHODS TO `Role`, CREATE
// ANOTHER .go FILE IN THE PACKAGE `role`

// Role: This is a Role description
package role

// RoleBase is the base struct with the default getters and setters
type RoleBase struct {
	id int64 // Primary key for role
	name string // Role name

	error error
	self *Role
}

// -------------------------------------------

// NewRole returns a pointer to a new NewRole object
func NewRole() *Role {
	role := &Role{}
	role.self = role
	return role
}

// -------------------------------------------

// Error return the error associate to the object
func (x *RoleBase) Error () error {
	return x.error
}

// -------------------------------------------

// Error return the error associate to the object
func (x *RoleBase) Validate () error {
	x.error = nil
	return x.error
}

// -------------------------------------------
// Id returns Role.id
func (x *RoleBase) Id () int64 {
	return x.id
}

// SetId sets Role.id
func (x *RoleBase) SetId (v int64) *Role {
	x.id = v
	return x.self
}
// Name returns Role.name
func (x *RoleBase) Name () string {
	return x.name
}

// SetName sets Role.name
func (x *RoleBase) SetName (v string) *Role {
	x.name = v
	return x.self
}

// -------------------------------------------

// String implements Stringer
func (x *RoleBase) String () string {
	jsonString, err := x.ToJSON()
	if err != nil {
		panic(err)
	}
	return "*Role" + jsonString
}
