package main

import (
	"log"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func ContentTypeHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	rw := negroni.NewResponseWriter(w)
	rw.Before(func(rw negroni.ResponseWriter) {
		if rw.Status() == http.StatusOK {
			rw.Header().Set("Content-Type", "application/json")
		}
	})
	next(rw, r)
}

func main() {

	databasePath := "/database/jackmarshall"
	database, err := db.OpenDB(databasePath)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	database.Create("Scenario")
	database.Create("Tournaments")
	// database, err := mgo.Dial(os.Getenv("DATABASE_PORT_27017_TCP_ADDR"))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer database.Close()

	router := httprouter.New()
	router.GET("/scenario", NewListScenarioHandler(database))
	// router.GET("/scenario/:id", NewGetScenarioHandler(database))
	router.POST("/scenario", NewCreateScenarioHandler(database))
	// router.PUT("/scenario/:id", NewUpdateScenarioHandler(database))
	// router.DELETE("/scenario/:id", NewDeleteScenarioHandler(database))

	// router.GET("/tables", NewListTableHandler(database))
	//
	// router.GET("/players", NewListPlayerHandler(database))
	//
	// router.GET("/tournaments/:id/round", NewCreateRoundHandler(database))
	//
	router.GET("/tournaments", NewListTournamentHandler(database))
	// router.GET("/tournaments/:id", NewGetTournamentHandler(database))
	router.POST("/tournaments", NewCreateTournamentHandler(database))
	// router.PUT("/tournaments/:id", NewUpdateTournamentHandler(database))
	// router.DELETE("/tournaments/:id", NewDeleteTournamentHandler(database))

	// Initialize the middleware stack
	stack := negroni.New()
	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())

	stack.Use(negroni.HandlerFunc(ContentTypeHandler))
	stack.UseHandler(router)

	log.Fatalln(http.ListenAndServe(":8081", stack))
}
