package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/HouzuoGuo/tiedot/data"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func main() {
	databasePath := "database"
	database, err := data.OpenCollection(databasePath)
	if err != nil {
		panic(err)
	}

	router := httprouter.New()
	router.GET("/api/tournaments", NewListTournamentHandler(database))
	router.GET("/api/tournaments/:id", NewGetTournamentHandler(database))
	router.POST("/api/tournaments", NewCreateTournamentHandler(database))
	router.PUT("/api/tournaments/:id", NewUpdateTournamentHandler(database))
	router.DELETE("/api/tournaments/:id", NewDeleteTournamentHandler(database))

	router.GET("/api/tournaments/:id/round", NewCreateRoundHandler(database))
	router.GET("/api/tournaments/:id/results", NewGetResultsHandler(database))

	router.NotFound = http.FileServer(http.Dir("front"))

	// Initialize the middleware stack
	stack := negroni.New()
	//	stack.Use(negroni.NewLogger())
	stack.Use(negroni.NewRecovery())
	stack.UseHandler(router)

	//testAssignement()
	createTestTournament(database, "testB&R - 8", 64, 32, 8)
	createTestTournament(database, "testB&R - 7", 64, 32, 7)
	createTestTournament(database, "testB&R - 6", 64, 32, 6)
	fmt.Println("Server running on localhost:8080...")
	log.Fatalln(http.ListenAndServe(":8080", stack))
}

func createTestTournament(database *data.Collection, name string, nbPlayers int, nbTables int, nbScenario int) {
	factions := []string{"Cygnar", "Trollbloods", "Legion of Everblight", "Cryx", "Khador"}
	tournament := Tournament{}
	players := []Player{}
	tables := []Table{}

	tournament.Name = name
	tournament.Rounds = []Round{}

	nbFaction := len(factions)
	lists := []List{List{Caster: "caster1"}, List{Caster: "caster2"}}
	for i := 0; i < nbPlayers; i++ {
		players = append(players, Player{
			Name:    "player" + fmt.Sprintf("%d", i),
			Faction: factions[rand.Intn(nbFaction)],
			Lists:   lists,
		})
	}

	for i := 0; i < nbTables; i++ {
		tables = append(tables, Table{Name: "table" + fmt.Sprintf("%d", i), Scenario: "scenario" + fmt.Sprintf("%d", i%nbScenario)})
	}

	tournament.Tables = tables
	tournament.Players = players

	data, err := json.Marshal(tournament)
	if err != nil {
		fmt.Println("ERROR MARSHALL")
	}

	_, err = database.Insert(data)
	if err != nil {
		fmt.Println("ERROR INSERT")

	}
}

func testAssignement() {
	tournament := Tournament{}
	players := []Player{}
	tables := []Table{}

	nbPlayers := 8
	nbTables := 4
	nbScenario := 4
	nbRounds := 2

	for i := 0; i < nbPlayers; i++ {
		players = append(players, Player{Name: "player" + fmt.Sprintf("%d", i)})
	}

	for i := 0; i < nbTables; i++ {
		tables = append(tables, Table{Name: "table" + fmt.Sprintf("%d", i), Scenario: "scenario" + fmt.Sprintf("%d", i%nbScenario)})
	}

	tournament.Tables = tables
	tournament.Players = players

	for i := 0; i < nbRounds; i++ {
		round := Round{
			Number: i,
			Games:  []Game{},
		}
		var pairings = CreatePairs(players, tournament, &round)

		createRound(pairings, tables, &round)

		for _, g := range round.Games {
			g.Results[rand.Intn(2)].VictoryPoints = 1
		}

		tournament.Rounds = append(tournament.Rounds, round)
	}
	displayTournament(tournament)
	fmt.Println("FINI ! ")
}

func displayTournament(tournament Tournament) {
	for _, round := range tournament.Rounds {
		fmt.Println(round.String())
	}
}
