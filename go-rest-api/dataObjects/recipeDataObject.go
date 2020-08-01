package dao

import (
	"log"

	. "go-rest-api/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RecipesDataObject struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "recipes"
)

func (rDO *RecipesDataObject) Connect() {
	session, err := mgo.Dial(rDO.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(rDO.Database)
}

func (rDO *RecipesDataObject) FindAll() ([]Recipe, error) {
	var recipes []Recipe
	err := db.C(COLLECTION).Find(bson.M{}).All(&recipes)
	return recipes, err
}

func (rDO *RecipesDataObject) Insert(recipe Recipe) error {
	err := db.C(COLLECTION).Insert(&recipe)
	return err
}

func (rDO *RecipesDataObject) FindById(id string) (Recipe, error) {
	var recipe Recipe
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&recipe)
	return recipe, err
}

func (rDO *RecipesDataObject) Update(id string, recipe Recipe) error {
	err := db.C(COLLECTION).Update(bson.M{"_id": bson.ObjectIdHex(id)}, &recipe)
	return err
}

func (rDO *RecipesDataObject) Delete(id string) error {
	err := db.C(COLLECTION).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (rDO *RecipesDataObject) RateRecipeById(id string, rating Rating) error {
	err := db.C(COLLECTION).Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$push": bson.M{"rating": &rating}})
	return err
}
