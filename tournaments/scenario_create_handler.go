package main

import (
	"encoding/json"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/julienschmidt/httprouter"
)

func NewCreateScenarioHandler(database *db.DB) httprouter.Handle {
	collection := database.Use("Scenario")
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		var scenario Scenario
		err := json.NewDecoder(r.Body).Decode(&scenario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = collection.Insert(map[string]interface{}{
			"name": scenario.Name,
			"year": scenario.Year,
			"link": scenario.Link})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// err = collection.Insert(&scenario)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		w.WriteHeader(http.StatusOK)
	}
}
