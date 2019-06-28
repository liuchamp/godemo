package main

import (
	"encoding/json"
	"fmt"
)

type RegisterDto struct {
	Name   string `form:"user" bson:"name,omitempty" json:"user" xml:"user"  binding:"required"`
	Passwd string `form:"password" bson:"password,omitempty" json:"password" xml:"password" binding:"required"`
}
type UserModel struct {
	AddUserDto

	BasicModels
}
type BasicModels struct {
	CreateTime int64 `json:"c_t" bson:"create_time,omitempty"`
	UpdateTime int64 `json:"u_t" bson:"update_time,omitempty"`
}

type UserCoreInfo struct {
	Phone        string `bson:"phone_number,omitempty" json:"phone_number" binding:"required"`
	WxAccunt     string `bson:"wx_account,omitempty" json:"wx_account" binding:"wxOptions"`
	RightOfLogin bool   `bson:"right_login,omitempty" json:"right_login" binding:"required"`
	Scope        string `bson:"scope,omitempty" json:"scope" binding:"required"`
}
type UserOptInfo struct {
	LastLoginTime int64  `bson:"last_login_time,omitempty" json:"last_login_time"`
	LastLoginIp   string `bson:"last_login_ip,omitempty" json:"last_login_ip"`
}
type AgentInfo struct {
	SettlePd    string `bson:"settle_password,omitempty" json:"settle_password"`
	AlipyAccunt string `bson:"aplipy_accunt,omitempty" json:"aplipy_accunt"`
	AlipyPerson string `bson:"aplipy_person,omitempty" json:"aplipy_person"`
	BankPerson  string `bson:"bank_person,omitempty" json:"bank_person"`
	BankName    string `bson:"bank_name,omitempty" json:"bank_name"`
	BankNumber  string `bson:"bank_number,omitempty" json:"bank_number"`
}

type AddUserDto struct {
	RegisterDto
	UserCoreInfo
}

func main() {
	rdto := RegisterDto{Name: "staidng", Passwd: "sdafdsaf"}
	urc := UserCoreInfo{Phone: "15946", RightOfLogin: true, Scope: "123,223"}
	addu := AddUserDto{}
	addu.UserCoreInfo = urc
	addu.RegisterDto = rdto
	bm := BasicModels{CreateTime: 1656056, UpdateTime: 1656056}
	um := UserModel{}
	um.AddUserDto = addu
	um.BasicModels = bm

	bs, _ := json.Marshal(um)

	fmt.Println(string(bs))

}
