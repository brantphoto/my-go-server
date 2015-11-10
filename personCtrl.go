package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

//ROOT CONTROLLERS
func PersonIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB(dbName).C("people")
	result := People{}
	c.Find(bson.M{"name": "Ale"}).All(&result)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func PersonShow(w http.ResponseWriter, r *http.Request) {

	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	c := sessionCopy.DB(dbName).C("people")
	result := Person{}
	err := c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
