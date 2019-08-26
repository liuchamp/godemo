package main

import (
	"fmt"
	"github.com/casbin/casbin"
	mongodbadapter "github.com/casbin/mongodb-adapter"
)

func main() {

	// Initialize a MongoDB adapter and use it in a Casbin enforcer:
	// The adapter will use the database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a := mongodbadapter.NewAdapter("192.168.0.193:27017") // Your MongoDB URL.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := mongodbadapter.NewAdapter("127.0.0.1:27017/abc")

	e, err := casbin.NewEnforcer("authz_model.conf", a)
	if err != nil {
		panic(err)
	}

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	s, err := e.Enforce("alice", "data1", "read")
	fmt.Println(s, err)

	e.AddPolicy("alice", "data1", "read")

	s, err = e.Enforce("alice", "data1", "read")
	fmt.Println(s, err)
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
