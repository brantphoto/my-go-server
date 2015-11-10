package main

import (
	"fmt"
	"net/http"
)

//ROOT CONTROLLERS
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

//TODO CONTROLLERS
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Show!")
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Create!")
}
