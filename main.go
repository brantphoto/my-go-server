package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

const dbName string = "test"
const dbUrl string = "bbs-comp.att.net:27017"

var session mgo.Session

func main() {
	session, err := mgo.Dial(dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	storeSession(session)

	//err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	//&Person{"Cla", "+55 53 8402 8510"})
	//if err != nil {
	//log.Fatal(err)
	//}

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func storeSession(s *mgo.Session) {
	session = *s
}
