package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
)

type Smou struct {
	SID  string `json:"_id" bson:"_id"`
	Scd  string `json:"scd" bson:"scd"`
	Smut string `json:"sd" bson:"sd"`
}
type TestDemo struct {
	ID  primitive.ObjectID `json:"_id" bson:"_id"`
	Sid string             `json:"sd" bson:"sd"`
	Sd  int64              `json:"ss" bson:"ss"`
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.0.196:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.TODO())

	client.Database("test").Collection("sdmo").DeleteMany(context.TODO(), bson.M{})
	//sduls()
	testdata(client)
}

func sduls() {
	input, err := ioutil.ReadFile("sdi.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var ss []TestDemo
	err = json.Unmarshal([]byte(input), &ss)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bs, _ := json.Marshal(ss)
	fmt.Println(string(bs))
}
func inster(client *mongo.Client) {
	input, err := ioutil.ReadFile("input.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var ss []Smou
	err = json.Unmarshal([]byte(input), &ss)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bs, err := json.Marshal(ss)
	fmt.Println(string(bs))
	var doc []interface{}
	for _, e := range ss {
		doc = append(doc, e)
	}

	sr, err := client.Database("test").Collection("sdmo").InsertMany(context.TODO(), doc)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range sr.InsertedIDs {
		fmt.Println(v)
	}
}

func testdata(client *mongo.Client) {
	doc := StatDaliyTotal{
		ID:                primitive.NewObjectID(),
		Date:              152020,
		DDate:             61651632165,
		ChanID:            "dsafidsnafids",
		Revenue:           999,
		Tax:               1,
		Win:               1,
		Recharge:          1,
		RecharNum:         1,
		PayRate:           1,
		AvgRecharge:       1,
		Exchange:          1,
		ExchangeNum:       1,
		ExchangeTax:       1,
		AvgExchange:       1,
		Player:            1,
		BindPlayer:        1,
		BindRate:          1,
		PlayerRecharge:    1,
		PlayerRechaRate:   1,
		AvgPlayerRecharge: 1,
		OldPlayerSignNum:  1,
		OldPlayerRecharge: 1,
		AvgOldPlayerRecha: 1,
		SignTotal:         1,
		CustomHandsel:     1,
		Zjh: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Ddz: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Lhd: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Qznn: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Brnn: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		By: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Hhdz: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Jdnn: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Dfdc: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Pdk: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Ermj: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Dzpk: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Hbsl: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
		Ebg: GameDayRecord{
			TotalTime: 1,
			RoundNum:  1,
			Tax:       1,
			Wol:       1,
		},
	}

	sr, err := client.Database("test").Collection("sdmo").InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sr.InsertedID)

}

type StatDaliyTotal struct {
	ID                primitive.ObjectID `bson:"_id"`               //记录id
	Date              int64              `bson:"time"`              //日期的unix值
	DDate             int64              `bson:"date"`              //计算日期格式是年月日
	ChanID            string             `bson:"chanID"`            //渠道id
	Revenue           int64              `bson:"revenue"`           //总营收
	Tax               int64              `bson:"tax"`               //总税收
	Win               int64              `bson:"wol"`               //总输赢
	Recharge          int64              `bson:"recharge"`          //总充值
	RecharNum         int64              `bson:"rechargeNum"`       //总充值人数
	PayRate           float64            `bson:"payRate"`           //付费率
	AvgRecharge       int64              `bson:"avgRecharge"`       //平均充值金额
	Exchange          int64              `bson:"exchage"`           //总兑换
	ExchangeNum       int64              `bson:"exchageNum"`        //总兑换人数
	ExchangeTax       int64              `bson:"exchangeTax"`       //总兑换手续费
	AvgExchange       int64              `bson:"agvExchange"`       //平均兑换金额
	Player            int64              `bson:"players"`           //新增用户
	BindPlayer        int64              `bson:"bindPlayer"`        //新增用户绑定数
	BindRate          float64            `bson:"bindRate"`          //新增用户绑定率
	PlayerRecharge    int64              `bson:"playerRecharge"`    //新用户充值金额
	PlayerRechaRate   float64            `bson:"playerRechaRate"`   //新用户付费率
	AvgPlayerRecharge int64              `bson:"avgPlayerRecharge"` //新增用户平均充值
	OldPlayerSignNum  int64              `bson:"oldPlayerSign"`     //老用户登录数
	OldPlayerRecharge int64              `bson:"oldPlayerRecharge"` //老用户充值金额
	AvgOldPlayerRecha int64              `bson:"avgOldPlayerRecha"` //老用户平均充值
	SignTotal         int64              `bson:"signTotal"`         //登录总人数, 表示总登陆人数
	CustomHandsel     int64              `bson:"customHandsel"`     //游客赠送
	Zjh               GameDayRecord      `bson:"zjh"`               //金花
	Ddz               GameDayRecord      `bson:"ddz"`               //斗地主
	Lhd               GameDayRecord      `bson:"lhd"`               //龙虎
	Qznn              GameDayRecord      `bson:"qznn"`              //抢庄牛牛
	Brnn              GameDayRecord      `bson:"brnn"`              //百人牛牛
	By                GameDayRecord      `bson:"by"`                //捕鱼
	Hhdz              GameDayRecord      `bson:"hhdz"`              //红黑
	Jdnn              GameDayRecord      `bson:"jdnn"`              //经典牛牛
	Dfdc              GameDayRecord      `bson:"dfdc"`              //多福多财
	Pdk               GameDayRecord      `bson:"pdk"`               //跑得快
	Ermj              GameDayRecord      `bson:"ermj"`              //二人麻将
	Dzpk              GameDayRecord      `bson:"dzpk"`              //德州扑克
	Hbsl              GameDayRecord      `bson:"hbsl"`              //红包扫雷
	Ebg               GameDayRecord      `bson:"ebg"`               //二八杠
}

//游戏天记录
type GameDayRecord struct {
	TotalTime int64 `bson:"totalTime"` //总时间
	RoundNum  int64 `bson:"roundNum"`  //局数
	Wol       int64 `bson:"wol"`
	Tax       int64 `bson:"tax"`
}
