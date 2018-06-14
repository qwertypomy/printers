package mongodb

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/qwertypomy/printers/config"
	"github.com/qwertypomy/printers/utils"
)

var Db *mgo.Database

func init() {
	cfg, err := config.GetConfig()
	utils.FatalError(err)
	Db = getSession().DB(cfg.Database)
}

func getSession() (session *mgo.Session) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	return
}
