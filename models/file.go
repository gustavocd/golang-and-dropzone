package models

import "gopkg.in/mgo.v2/bson"

// File represents file's information.
type File struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Path string        `json:"path" bson:"path"`
}

// Create saves a file into the db
func (f File) Create() (err error) {
	connection, _ := getConnection()
	err = connection.DB("files").C("images").Insert(f)
	if err != nil {
		return
	}
	return nil
}
