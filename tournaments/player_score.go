package main

type PlayerScoreList []PlayerScore

type PlayerScore struct {
	Player string `json:"player"`
	Score  int    `json:"score"`
}

func (p PlayerScoreList) Len() int           { return len(p) }
func (p PlayerScoreList) Less(i, j int) bool { return p[i].Score < p[j].Score }
func (p PlayerScoreList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
