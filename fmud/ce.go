package fmud

import "fmt"

type Inum struct {
}

func NewInum() *Inum {
	return &Inum{}
}

func (sd Inum) Find() {
	fmt.Println("Find")
}
func (sd Inum) Geter() {
	fmt.Println("Geter")
}

type InUs struct {
	Inum
}

func NewInUs() *InUs {
	return &InUs{}
}
