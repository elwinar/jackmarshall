package main

import (
	"encoding/json"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/julienschmidt/httprouter"
)

func NewListScenarioHandler(database *db.DB) httprouter.Handle {
	collection := database.Use("Scenario")

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		results := []Scenario{}

		collection.ForEachDoc(func(id int, doc []byte) (willMoveOn bool) {
			var scenario Scenario
			err := json.Unmarshal(doc, &scenario)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			results = append(results, scenario)
			return true // move on to the next document OR
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}
