package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"golang.org/x/tools/go/ast/astutil"
)

func main() {
	demo()
}

func demo2() {
	fset := token.NewFileSet()
	fss, err := parser.ParseDir(fset, "/Users/mac/Code/go/src/github.com/liuchamp/godemo/models", sdik, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(fss))
	fssd, err := parser.ParseDir(fset, "/Users/mac/Code/go/src/github.com/liuchamp/godemo/models/redels", sdik, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(fssd))

	ast.Print(fset, fss)

}

func sdik(info os.FileInfo) bool {
	return true
}

func demo() {
	// 这就是上一章的代码.
	src := `
package main
import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)
func main() {
    println("Hello, World!")
}
`
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset

	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	// astutil.AddImport(fset, f, "app")
	// Print the AST.
	for _, v := range f.Imports {
		fmt.Println(v.Path.Value)
	}

	sfm := astutil.Imports(fset, f)
	for _, v := range sfm {
		for _, s := range v {
			fmt.Println(s.Path.Value)
		}
	}
	ast.Print(fset, f)

}
