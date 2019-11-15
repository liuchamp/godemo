package model

import "github.com/vmihailenco/msgpack"

type User struct {
	Username string `json:"username"`
	RoleID   string `json:"roleId"`
}

func (s *User) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(s)
}

// UnmarshalBinary use msgpack
func (s *User) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, s)
}
