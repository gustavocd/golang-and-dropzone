package models

import (
	"time"

	"gopkg.in/mgo.v2"
)

// getConnection gets a session of mongodb
func getConnection() (session *mgo.Session, err error) {
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Database: "files",
		Username: "",
		Password: "",
		Timeout:  60 * time.Second,
	}

	session, err = mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		return nil, err
	}
	defer session.Clone()
	session.SetMode(mgo.Monotonic, true)

	if session != nil {
		return
	}
	return
}
