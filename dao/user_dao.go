package dao

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/binarylibrarian/escapement/models"
	"gopkg.in/mgo.v2"
)

// UserDAO manages database interactions for the User model
type UserDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION is the mongodb "collection" for this doa
	COLLECTION = "users"
)

func (udao *UserDAO) connect() {
	session, err := mgo.Dial(udao.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(udao.Database)
}

// Insert addeds the user to the "users" mongodb collection
func (udao *UserDAO) Insert(u models.User) error {
	err := db.C(COLLECTION).Insert(&u)
	return err
}

// FindByID returns a user from the mongodb "users" collection
func (udao *UserDAO) FindByID(id string) (models.User, error) {
	var user models.User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)

	return user, err
}
