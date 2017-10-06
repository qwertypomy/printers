package mongodb

import (
	"gopkg.in/mgo.v2"
)

var Db *mgo.Session

func init() {
	Db = getSession()
}

func getSession() (session *mgo.Session) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	return
}
