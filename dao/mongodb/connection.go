package mongodb

//import (
//	"database/sql"
//	"gopkg.in/mgo.v2"
//)
//
//var Db *sql.DB
//
//func init() {
//	Db = getDB()
//}
//
//func getDB() (session *mgo.Session, err error) {
//	session, err = mgo.Dial("localhost")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//
//	session.SetMode(mgo.Monotonic, true)
//	ensureIndex(session)
//	return
//}
//
//func ensureIndex(s *mgo.Session) {
//	session := s.Copy()
//	defer session.Close()
//
//	c := session.DB("store").C("books")
//
//	index := mgo.Index{
//		Key:        []string{"isbn"},
//		Unique:     true,
//		DropDups:   true,
//		Background: true,
//		Sparse:     true,
//	}
//	err := c.EnsureIndex(index)
//	if err != nil {
//		panic(err)
//	}
//}
