package main

import "gopkg.in/mgo.v2"

type DataStore struct {
	session mgo.Session
}
