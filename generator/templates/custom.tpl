// {{ .Name | fc }}: {{ .Description }}
package {{ .Name | lc }}


// {{ .Name | fc }} allows to overwrite the base class
type {{ .Name | fc }} struct {
	{{ .Name | fc }}Base 

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
