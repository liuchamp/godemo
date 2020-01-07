package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

type PayInfoModel struct {
	CardHolderName string `json:"cardHolderName" bson:"cardHolderName"`
	AlipayName     string `json:"alipayName" bson:"alipayName"`
	CardNo         string `json:"cardNo" bson:"cardNo"`
	CardType       string `json:"cardType" bson:"cardType"`
	AlipayAccount  string `json:"alipayAccount" bson:"alipayAccount"`
}

func main() {
	bs, err := json.Marshal(PayInfoModel{
		CardHolderName: "start_user",
		AlipayName:     "start_user",
		CardNo:         "6241941651496819856",
		CardType:       "中国银行",
		AlipayAccount:  "ali1561651564156",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bs))

	encodeString := base64.StdEncoding.EncodeToString(bs)
	log.Println(encodeString)

	// 如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString([]byte(bs))
	log.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(uDec))

}
