package main

import (
	"io"
	"os"
	"sort"
)

type TeamName string

type Team struct {
	Name        TeamName
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[TeamName]int
}

type Score int

// this works because maps are implemented using pointers
// so even though l is copied when passed, l.Wins is a pointer, and the pointer value (memory address) stays the same
func (l League) MatchResult(tn1 TeamName, s1 Score, tn2 TeamName, s2 Score) {
	if _, ok := l.Wins[tn1]; !ok {
		return
	}
	if _, ok := l.Wins[tn2]; !ok {
		return
	}
	if s1 > s2 {
		l.Wins[tn1]++
	}
	if s2 > s1 {
		l.Wins[tn2]++
	}
}

type Ranker interface {
	Ranking() []TeamName
}

func (l League) Ranking() []TeamName {
	teamNames := make([]TeamName, 0, len(l.Teams))
	for tn := range l.Wins {
		teamNames = append(teamNames, tn)
	}

	sort.Slice(teamNames, func(i, j int) bool {
		return l.Wins[teamNames[i]] > l.Wins[teamNames[j]]
	})

	return teamNames
}

func RankPrinter(r Ranker, wr io.Writer) {
	for _, tn := range r.Ranking() {
		io.WriteString(wr, string(tn)+"\n")
	}
}

func Ch7() {
	team1 := Team{Name: "Team 1", PlayerNames: []string{"Bill", "Bob"}}
	team2 := Team{Name: "Team 2", PlayerNames: []string{"John", "Jane"}}
	team3 := Team{Name: "Team 3", PlayerNames: []string{"Boris", "Rishi"}}

	league := League{Teams: []Team{team1, team2, team3}, Wins: map[TeamName]int{}}
	league.Wins[team1.Name] = 0
	league.Wins[team2.Name] = 0
	league.Wins[team3.Name] = 0

	league.MatchResult(team1.Name, 5, team2.Name, 1)
	league.MatchResult(team3.Name, 4, team2.Name, 3)
	league.MatchResult(team3.Name, 3, team1.Name, 1)

	RankPrinter(league, os.Stdout)
}
