package main

import (
	"fmt"
)

func main() {
	//src, err := ioutil.ReadFile("pikm/test.go")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fset := token.NewFileSet()
	//file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	//if err != nil {
	//	panic(err)
	//}
	//
	//ast.Inspect(file, func(x ast.Node) bool {
	//	s, ok := x.(*ast.StructType)
	//	if !ok {
	//		return true
	//	}
	//
	//	for _, field := range s.Fields.List {
	//		fmt.Printf("Field: %s\n", field.Names[0].Name)
	//		fmt.Printf("Tag:   %s\n", field.Tag.Value)
	//	}
	//	return false
	//})
	//

	fmt.Println(len("5e5e107fb351a231750701f4"))
	fmt.Println(len("ff4ecde7-826d-4660-985b-2b1354297971") * 100000)

}
