/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"encoding/json"
	"fmt"
)

type PlayerRealFlow struct {
	TotalRecharge int64 `json:"totalRecharge" bson:"totalRecharge"` //总充值(单位元)
	ScoreFlow     int64 `json:"scoreFlow" bson:"scoreFlow"`         //总流水
	ByFlow        int64 `json:"byFlow" bson:"byFlow"`               //捕鱼流水
	LhjFlow       int64 `json:"lhjFlow" bson:"lhjFlow"`             //老虎机流水
}

func main() {
	mk := PlayerRealFlow{
		TotalRecharge: 10000,
		ScoreFlow:     123,
		ByFlow:        1234123412,
		LhjFlow:       1234123412314,
	}
	str, _ := json.Marshal(mk)

	dmk := string(str)
	fmt.Println(dmk)
	data := &PlayerRealFlow{}
	if err := json.Unmarshal([]byte(dmk), data); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v", data)
	if err := json.Unmarshal([]byte("{\"totalRecharge\":100000000,\"scoreFlow\":120000000,\"byFlow\":1012000000,\"lhjFlow\":1560000000}"), data); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v", data)

}
