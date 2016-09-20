package main

import (
	"encoding/json"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/julienschmidt/httprouter"
)

func NewListTournamentHandler(database *db.DB) httprouter.Handle {
	collection := database.Use("Tournaments")

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		results := []Tournament{}

		collection.ForEachDoc(func(id int, doc []byte) (willMoveOn bool) {
			var tournament Tournament
			err := json.Unmarshal(doc, &tournament)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			results = append(results, tournament)
			return true // move on to the next document OR
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(results)
	}
}
