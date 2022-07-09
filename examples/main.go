package main

import (
	"fmt"

	"github.com/frncscsrcc/lazygen/examples/model/domain/role"
	"github.com/frncscsrcc/lazygen/examples/model/domain/user"
)

func main() {
	role1 := role.NewRole().SetId(1).SetName("RoleName")

	user1 := user.NewUser().
		SetId(2).
		SetUsername("UsernameTest").
		SetName("MyName").
		SetRole(role1).
		AppendPhone("PHONE1").
		AppendPhone("PHONE2").
		SetPassword("MyPasswordPlain")

	fmt.Printf("%v\n", user1)
}
