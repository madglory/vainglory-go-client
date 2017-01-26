package main

import (
	"fmt"

	v "github.com/madglory/vainglory-go-client"
)

var apikey = "aaa.bbb.ccc"
var matchID = "f78917d2-d7cf-11e6-ad79-062445d3d668"
var playerID = "3a8f6440-1b20-11e6-a24a-06ab1a16f8e5"
var offset = "0"
var limit = "3"

func main() {

	q := new(v.QueryRequest)

	q.Limit = limit
	q.Offset = offset

	//Creates the client
	client := v.NewClient(apikey, q)

	//Return an object, serverErrorResponse, and other Error
	match, _, _ := client.GetMatchByID(matchID)
	player, _, _ := client.GetPlayerByID(playerID)

	//Returns a string, serverErrorResponse, and other Error
	matches, _, _ := client.GetMatches()

	//Return a matches object
	fmt.Printf("3 Matches: \n%v\n", matches)
	fmt.Printf("Game mode of single match: %v\n", match.GameMode)
	fmt.Printf("Single player name: %v\n", player.Name)

}
