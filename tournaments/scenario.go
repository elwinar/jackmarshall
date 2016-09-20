package main

// import (
// 	"gopkg.in/mgo.v2/bson"
// )

type Scenario struct {
	//	ID   bson.ObjectId `json:"id"`
	Id   string `json:"id"`
	Name string `json:"name"`
	Year int    `json:"year"`
	Link string `json:"link"`
}
