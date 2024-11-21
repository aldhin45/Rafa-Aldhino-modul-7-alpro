package main

import (
	"fmt"
)

type Match struct {
	TeamA  string
	TeamB  string
	ScoreA int
	ScoreB int
	Winner string
}

func main() {
	var matches []Match
	var teamA, teamB string
	var scoreA, scoreB int

	fmt.Println("Masukkan nama Klub A dan Klub B:")
	fmt.Scan(&teamA, &teamB)

	for {
		fmt.Printf("Masukkan skor pertandingan (%s vs %s): ", teamA, teamB)
		fmt.Scan(&scoreA, &scoreB)
		if scoreA < 0 || scoreB < 0 {
			break
		}

		var winner string
		if scoreA > scoreB {
			winner = teamA
		} else if scoreB > scoreA {
			winner = teamB
		} else {
			winner = "Draw"
		}

		matches = append(matches, Match{
			TeamA:  teamA,
			TeamB:  teamB,
			ScoreA: scoreA,
			ScoreB: scoreB,
			Winner: winner,
		})
	}

	fmt.Println("\nHasil pertandingan:")
	for i, match := range matches {
		fmt.Printf("Pertandingan %d: %s vs %s | Skor: %d - %d | Pemenang: %s\n",
			i+1, match.TeamA, match.TeamB, match.ScoreA, match.ScoreB, match.Winner)
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for _, match := range matches {
		if match.Winner != "Draw" {
			fmt.Println(match.Winner)
		}
	}
}
