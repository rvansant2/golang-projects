package models

import "gopkg.in/mgo.v2/bson"

type Recipe struct {
	UniqueID   bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string        `bson:"name" json:"name"`
	PrepTime   string        `bson:"prep_time" json:"prep_time"`
	Difficulty string        `bson:"difficulty" json:"difficulty"`
	Vegetarian string        `bson:"vegetarian" json:"vegetarian"`
	Rating     []Rating      `bson:"rating"`
}

type Rating struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Rating string        `bson:"rating,omitempty" json:"rating,omitempty"`
}
