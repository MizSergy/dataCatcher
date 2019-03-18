package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

var session *mgo.Session
var database *mgo.Database

func MongoConnect() {
	var err error
	session, err = mgo.Dial(os.Getenv("MONGO_HOSTS"))
	if err != nil {
		panic(err.Error())
	}

	err = session.Login(&mgo.Credential{
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})
	if err != nil {
		panic(err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
}

func MongoSession() *mgo.Session {
	return session
}
func MongoDB() *mgo.Database {
	if database != nil {
		return database
	}
	confDb := os.Getenv("MONGO_DB")
	database = session.DB(confDb)
	return database
}
func GetValueForNextSequence(name string) int {
	counter := MongoDB().C("counter")

	var result struct {
		Id            int "_id"
		SequenceValue int "sequence_value"
	}

	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"sequence_value": 1}},
		ReturnNew: true,
	}
	_, err := counter.Find(bson.M{"_id": name}).Apply(change, &result)

	if err != nil {
		//panic(err)
		counter.Insert(bson.M{"_id": name, "sequence_value": 1})
		return 1
	}
	return result.SequenceValue
}

func MongoDissconnect() {
	session.Close()
}
