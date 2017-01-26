package vainglory

import "fmt"

func PrettyPlayer(p *Player) {

	fmt.Println("\nPlayer Data:" + "\nID: " + p.ID + "\nName: " + p.Name + "\nTitleID: " +
		p.TitleID + "\nShardID: " + p.ShardID + "\nStats: ")

	for k := range p.Stats {
		fmt.Printf("%v : %v\n", k, p.Stats[k])
	}
}
