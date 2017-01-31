package vainglory

import "testing"

var apikey = "aaa.bbb.ccc"
var offset = "0"
var limit = "3"

func TestGetStatus(t *testing.T) {

	q := new(QueryRequest)
	client := NewClient(apikey, q)

	_, serverResponse, err := client.GetStatus()

	if err != nil {
		t.Fatalf("There was an error getting status: %v\n", err)
	}

	if serverResponse != 0 {
		t.Fatalf("Invalid server response: %v\n", serverResponse)
	}
}

func TestGetMatches(t *testing.T) {

	q := new(QueryRequest)
	q.Limit = "3"
	q.Offset = "0"
	q.SortField = "createdAt"

	client := NewClient(apikey, q)

	matches, serverResponse, err := client.GetMatches()

	if err != nil {
		t.Fatalf("The following error was thrown: %v", err)
	}
	if serverResponse != 0 {
		t.Fatalf("Invalid server response: %v\n", serverResponse)
	}
	if matches == nil {
		t.Fatalf("No matches found. ")
	}
}

func TestGetMatchByID(t *testing.T) {

	q := new(QueryRequest)
	client := NewClient(apikey, q)

	match, serverResponse, err := client.GetMatchByID("f78917d2-d7cf-11e6-ad79-062445d3d668")
	if err != nil {
		t.Fatalf("There was an error getting your match %v", err)
	}

	if match.ID != "f78917d2-d7cf-11e6-ad79-062445d3d668" {
		t.Fatalf("Got the wrong match")
	}

	if serverResponse != 0 {
		t.Fatalf("Invalid server response: %v\n", serverResponse)
	}
}

func TestGetPlayerByID(t *testing.T) {

	q := new(QueryRequest)
	client := NewClient(apikey, q)

	player, serverResponse, err := client.GetPlayerByID("3a8f6440-1b20-11e6-a24a-06ab1a16f8e5")
	if err != nil {
		t.Fatalf("There was an error getting your player %v", err)
	}

	if player.ID != "3a8f6440-1b20-11e6-a24a-06ab1a16f8e5" {
		t.Fatalf("Got the wrong player")
	}

	if serverResponse != 0 {
		t.Fatalf("Invalid server response: %v\n", serverResponse)
	}
}
