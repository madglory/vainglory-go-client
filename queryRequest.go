package vainglory

//QueryRequest exported
type QueryRequest struct {
	ShardID         string
	Offset          string
	Limit           string
	MatchID         string
	SortField       string
	PlayerNames     string
	TeamNames       string
	CreatedAtStart  string
	CreatedAtEnd    string
	ContentEncoding string
}
