package main

import (
	"fmt"
	"github.com/liuchamp/godemo/post"
)

func main() {
	p := &post.Post{}
	pkg := p.GetPackage()
	fmt.Println(pkg)
}
