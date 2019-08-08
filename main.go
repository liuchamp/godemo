package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j := `{
    "Name": "王老五",
    "Oks": "dsadsfa"
}`
	var jm map[string]interface{}
	json.Unmarshal([]byte(j), &jm)
	fmt.Print(jm)
}
