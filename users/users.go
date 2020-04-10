package users

import (
	"fmt"
	"github.com/liuchamp/godemo/person"
)

type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) error {
	fmt.Println(id)
	return u.Person.Get(id)
}
