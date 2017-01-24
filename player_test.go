package vainglory

func prettyPlayer(p *Player) string {

	return "\nID: " + p.ID + "\nName: " + p.Name + "\nTitleID: " +
		p.TitleID + "\nShardID: " + p.ShardID

}
