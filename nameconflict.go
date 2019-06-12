package main

// 名称冲突问题

type X struct {
	Name string
}
type Y struct {
	X
	Name string
}
