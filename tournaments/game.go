package main

type Game struct {
	Table   string   `json:"table"`
	Pairing []Result `json:"pairing"`
}
