package main

import (
	"flowershop/backend/internal/data"
	"fmt"
	"net/http"
	"time"
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

	flower := data.Flower{
		ID: id,
		CreatedAt: time.Now(),
		Name: "Lily",
		Origin: "",
		Details: "Lilium is a genus of herbaceous flowering plants growing from bulbs, all with large and often prominent flowers. Lilies are a group of flowering plants which are important in culture and literature in much of the world.",
		Price: 12,
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"flower": flower}, nil) 
	if err != nil{
		app.logger.Error(err.Error())
		http.Error(w, "server encountered problem" ,http.StatusInternalServerError)
	}

}