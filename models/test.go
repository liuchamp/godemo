package models

type Test struct {
	Name string
	Oks  string
	Smod Sim
	Next *Test `json:"next,omitempty"`
}

type Sim struct {
	Name string
	Mod  string
}
