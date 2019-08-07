package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	DemosMap()
}

func DemosMap() {

	jsonString := `{
    "b": "2",
    "a": "1",
    "c": "3"
}`
	var oo map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &oo)
	if err != nil {
		panic(err)
	}
	for k, v := range oo {
		fmt.Println(k, v)
	}

	jss := `["string","json","extent"]`
	var sars []string
	err = json.Unmarshal([]byte(jss), &sars)
	if err != nil {
		panic(err)
	}
	for k, v := range sars {
		fmt.Println(k, v)
	}
}
