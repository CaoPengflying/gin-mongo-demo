package entity

import "time"

type Vip struct {
	Name        string        `json:"name" bson:"name"`
	Mobile      string        `json:"mobile" bson:"mobile"`
	VipNo       string        `json:"vip_no" bson:"vip_no"`
	Integral    Integral      `json:"integral" bson:"integral"`
	ChangeFlows []ChangeFlow `json:"change_flows,omitempty" bson:"change_flows,omitempty"`
	Areas       []Area       `json:"areas" bson:"areas"`
}
type Integral struct {
	LeftInte     int       `bson:"left_inte"`
	SumInte      int       `bson:"sum_inte"`
	DeadlineDate time.Time `bson:"deadline_date"`
}

//changeFlow
type ChangeFlow struct {
	ChangeTime time.Time `json:"change_time" bson:"deadline_date"`
	ChangeInfo string    `json:"change_info" bson:"change_info"`
}

type Area struct {
	Province string `bson:"province"`
	City     string `bson:"city"`
	Address  string `bson:"address"`
}
