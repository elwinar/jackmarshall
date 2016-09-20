package main

import "gopkg.in/mgo.v2/bson"

type Tournament struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Owner     int           `json:"owner" bson:"owner"`
	Name      string        `json:"name"`
	Format    int           `json:"format"`
	Slots     int           `json:"slots"`
	FeeAmount float32       `json:"fee_amount"`
	//	Date      time.Time     `json:"date"`
	Players []Player `json:"players"`
	Tables  []Table  `json:"tables"`
	Rounds  []Round  `json:"rounds"`
}

func (t Tournament) isPairingOk(player1 string, player2 string) bool {
	for _, r := range t.Rounds {
		for _, g := range r.Games {
			if (g.Pairing[0].Player == player1 && g.Pairing[1].Player == player2) || (g.Pairing[0].Player == player2 && g.Pairing[1].Player == player1) {
				return false
			}
		}
	}
	return true
}

func (t Tournament) getVictoryPoints(player string) (res int) {
	for _, r := range t.Rounds {
		for _, g := range r.Games {
			for _, p := range g.Pairing {
				if p.Player == player {
					res += p.VictoryPoints
				}
			}
		}
	}
	return
}

func (t Tournament) getTablesPlayed(player string) (tables []string) {
	for _, r := range t.Rounds {
		for _, g := range r.Games {
			if g.Pairing[0].Player == player || g.Pairing[1].Player == player {
				tables = append(tables, g.Table)
			}
		}
	}
	return
}
