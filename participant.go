package vainglory

//Participant structure
type Participant struct {
	ID     string                 `json:"id" jsonapi:"primary,participant"`
	Player *Player                `json:"player" jsonapi:"relation,player"`
	Stats  map[string]interface{} `json:"stats" jsonapi:"attr,stats"`
	Actor  string                 `json:"actor" jsonapi:"attr,actor"`
}
