package entity

import "time"

type Vip struct {
	Name        string       `bson:"name"`
	Mobile      string       `bson:"mobile"`
	VipNo       string       `bson:"vip_no"`
	Integral    Integral     `bson:"integral"`
	ChangeFlows []ChangeFlow `bson:changeFlows`
	Areas       []Area       `bson:areas`
}
type Integral struct {
	LeftInte     int       `bson:"left_inte"`
	SumInte      int       `bson:"sum_inte"`
	DeadlineDate time.Time `bson:"deadline_date"`
}

type ChangeFlow struct {
	ChangeTime time.Time `bson:"deadline_date"`
	ChangeInfo string    `bson:"change_info"`
}

type Area struct {
	Province string `bson:"province"`
	City     string `bson:"city"`
	Address  string `bson:"address"`
}
