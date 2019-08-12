package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	v := validator.New()
	validateVariable(v)
}

func validateVariable(validate *validator.Validate) {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}
	myEmail2 := "dsafdsiangsd@gmail.com"
	errs = validate.Var(myEmail2, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}
	// email ok, move on
}
