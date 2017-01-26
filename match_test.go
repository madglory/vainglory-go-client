package vainglory

import "fmt"

func PrettyMatch(match *Match) string {

	fmt.Println("Participants data")
	x := 0
	y := 0
	for x < 2 {
		for y < 3 {
			prettyParticipant(match.Rosters[x].Participants[y])
			y++
		}
		x++
		y = 0
	}

	//Prints out the match rosters information
	r := match.Rosters
	fmt.Printf("Rosters: \n")
	for w := range r {
		fmt.Printf("%v : %v\n", w, r[w])
	}

	//prints out the rest of mactch stats
	s := match.Stats
	for k := range s {
		fmt.Printf("%v : %v\n", k, s[k])
	}
	return "\nMatch ID: " + match.ID + "\nGame mode: " + match.GameMode +
		"\nShard ID: " + match.ShardID + "\nTitleID: " + match.TitleID
}

func prettyParticipant(p *Participant) {
	fmt.Printf("Player Name: %v \n", p.Player.Name)
	fmt.Printf("-----------------------------\n")
	x := p.Stats
	for i := range x {
		fmt.Printf("%v : %v\n", i, x[i])
	}
	fmt.Println()
}
