package main

type Result struct {
	Player            string `json:"player"`
	List              int    `json:"list"`
	VictoryPoints     int    `json:"victory_points"`
	ScenarioPoints    int    `json:"scenario_points"`
	DestructionPoints int    `json:"destruction_points"`
	Buy               bool   `json:"buy"`
}
