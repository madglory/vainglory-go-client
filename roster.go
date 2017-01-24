package vainglory

//Roster structure
type Roster struct {
	ID           string                 `json:"id" jsonapi:"primary,roster"`
	Team         *Team                  `json:"team" jsonapi:"relation,team"`
	Participants []*Participant         `json:"participants" jsonapi:"relation,participants"`
	Stats        map[string]interface{} `json:"stats" jsonapi:"attr,stats"`
	Winner       string                 `json:"won"`
}
