package pikm

//db_start 刚开始
//发现多个注解
type AgentGameItem struct { //ssssss
	RoundNum     uint32  `json:"roundNum" bson:"roundNum"`         //局数
	WinScore     int64   `json:"winScore" bson:"winScore"`         //游戏消耗分数(输赢情况)
	Tax          int64   `json:"tax" bson:"tax"`                   //税收
	Flow         int64   `json:"flow" bson:"flow"`                 //流水
	InvisibleTax int64   `json:"InvisibleTax" bson:"InvisibleTax"` //暗税
	OnlinePlayer int64   `json:"onlinePlayer" bson:"onlinePlayer"` // 在线人数
	Percentage   float32 `json:"percentage" bson:"percentage"`     //占百分比（game是type时，表示各个游戏间的百分比，game时kind时，各个场次的百分比）
}

//db_start 刚开始
//发现多个注解
type AgentGamesItem struct {
	RoundNum uint32 `json:"roundNum" bson:"roundNum"` //局数
	//发现新大陆
	WinScore     int64   `json:"winScore" bson:"winScore"`         //游戏消耗分数(输赢情况)
	Tax          int64   `json:"tax" bson:"tax"`                   //税收
	Flow         int64   `json:"flow" bson:"flow"`                 //流水
	InvisibleTax int64   `json:"InvisibleTax" bson:"InvisibleTax"` //暗税
	OnlinePlayer int64   `json:"onlinePlayer" bson:"onlinePlayer"` // 在线人数
	Percentage   float32 `json:"percentage" bson:"percentage"`     //占百分比（game是type时，表示各个游戏间的百分比，game时kind时，各个场次的百分比）
}
