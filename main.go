package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

func main() {
	src, err := ioutil.ReadFile("pikm/test.go")
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(file, func(x ast.Node) bool {
		s, ok := x.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range s.Fields.List {
			fmt.Printf("Field: %s\n", field.Names[0].Name)
			fmt.Printf("Tag:   %s\n", field.Tag.Value)
		}
		return false
	})
}
