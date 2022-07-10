package main

import (
	"fmt"

	"github.com/frncscsrcc/lazygen/examples/model/domain/address"
	"github.com/frncscsrcc/lazygen/examples/model/domain/role"
	"github.com/frncscsrcc/lazygen/examples/model/domain/user"
)

func main() {
	fmt.Print("SERIALIZATION\n----------------------------------\n")
	role1 := role.NewRole().SetId(1).SetName("RoleName")

	user1 := user.NewUser().
		SetId(2).
		SetUsername("UsernameTest").
		SetName("MyName").
		SetRole(role1).
		AppendPhone("PHONE1").
		AppendPhone("PHONE2").
		SetPassword("MyPasswordPlain").
		AppendAddresses(
			address.NewAddress().
				SetStreet("street_a").
				SetNumber("1A").
				SetZip("11111").
				SetCountry("Germany"),
		).
		AppendAddresses(
			address.NewAddress().
				SetStreet("street_b").
				SetNumber("2A").
				SetZip("22222").
				SetCountry("Italy"),
		)

	user1Json, err := user1.ToJSON()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n\n", user1Json)

	fmt.Print("DESERIALIZATION\n----------------------------------\n")

	user2Json := `{"username": "myUsername", "addresses": [{"street":"street1"},{"street":"street2"}]}`
	user2, err := user.FromJSON(user2Json)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user2)
}
