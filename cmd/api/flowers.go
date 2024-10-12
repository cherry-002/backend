package main

import (
	"fmt"
	"net/http"
)

func (app *application) addFlowerHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "create a new flower")
}

func (app *application) showFlowerHandler(w http.ResponseWriter, r *http.Request) {


	id, err := app.readIDparam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of my flower %d\n", id)
}