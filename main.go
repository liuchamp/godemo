package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"text/template"
)

func main() {
	parseText()
}
func readFile() error {
	in, e := os.Open("main.go")
	if e != nil {
		return e
	}
	defer in.Close()
	inputReader := bufio.NewReader(in)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Println(inputString)
		if readerError == io.EOF {
			return readerError
		}
	}
	return nil
}
func parseGofile() error {

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "models/user.go", nil, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	// ast.Print(fset, f)

	for _, i := range f.Comments {
		log.Printf("comment:	%s", i.Text())
	}

	for _, i := range f.Decls {

		if fn, ok := i.(*ast.FuncDecl); ok {
			log.Printf("function:	%s", fn.Name.Name)
		}
		if fn, ok := i.(*ast.GenDecl); ok {
			log.Printf("gen:	%s", fn.Lparen)
		}

	}

	return nil
}

var html = `
"test").Parse("{{.Count}} items are made of {{.Material}}"
	{{$a := .Count}}
	{{$b := 17}}
	{{$c := 18}}	
  
	{{if eq  .Count $b}}
	oo
	{{else}}
xx
	{{end}}

	`

func parseText() {
	type Inventory struct {
		Material string
		Count    int
	}
	sweaters := Inventory{"axe", 0}

	tmpl, err := template.New("test").Parse(html)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

type Friend struct {
	Fname string
}

type Person struct {
	Username string
	Emails   []string
	Friends  []*Friend
}

func umsk() {
	f1 := Friend{Fname: "dsadignadf"}
	f2 := Friend{Fname: "dsadsafdsafdsdignadf"}
	p := Person{
		Username: "dsaidsnadsfa",
		Emails:   []string{"dsaidngaf", "dsagdsafdsa"},
		Friends:  []*Friend{&f1, &f2},
	}
	t := template.New("test")
	t = template.Must(t.Parse(
		`hello {{.Username}} ! 
{{ range .Emails}}
an email {{.}}
{{- end}}
{{with .Friends}}
{{- range .}}
friend name is {{.Fname}}
{{- end}}
{{end}}
`,
	))
	t.Execute(os.Stdout, p)
}

type EntetiesClass struct {
	Name  string
	Value int32
}

// In the template, we use rangeStruct to turn our struct values
// into a slice we can iterate over
const htmlTemplate = `
{{range $index, $element := .}}
    {{$index}}
        {{range $element}}
            {{.Name}} {{.Value}}
        {{end}}
{{end}}
`

func stmul() {
	data := map[string][]EntetiesClass{
		"Yoga":    {{"Yoga1", 13}, {"Yoga2", 15}},
		"Pilates": {{"Pilates1", 3}, {"Pilates2", 6}, {"Pilates3", 9}},
	}

	t := template.New("t")
	t, err := t.Parse(htmlTemplate)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
