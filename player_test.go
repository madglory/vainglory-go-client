package vainglory

import "fmt"

func prettyPlayer(p *Player) string {

	m := p.Stats

	for k := range m {
		fmt.Printf("%v : %v\n", k, m[k])
	}

	return "\nID: " + p.ID + "\nName: " + p.Name + "\nTitleID: " +
		p.TitleID + "\nShardID: " + p.ShardID

}
