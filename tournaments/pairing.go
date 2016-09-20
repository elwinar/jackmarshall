package main

import "fmt"

type Pairing struct {
	Player1 string `json:"id"`
	Player2 string `json:"id"`
}

func CreatePairings(players []string, tournament Tournament) (pairings []Pairing) {
	//for _, p1 := range players {
	for len(players) > 0 {
		for i2, p2 := range players {
			if players[0] == p2 {
				continue
			}
			if tournament.isPairingOk(players[0], p2) {
				fmt.Println("ma liste : ", players)
				fmt.Println("je paire " + players[0] + " et " + p2)
				pairings = append(pairings, Pairing{players[0], p2})
				players = append(players[:i2], players[i2+1:]...)
				fmt.Println("je retire "+p2+" de la liste, il reste : ", players)
				players = append(players[:0], players[1:]...)
				fmt.Println("je retire X de la liste, il reste : ", players)
				break
			}
		}
	}
	return
}
