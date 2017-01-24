package vainglory

import "testing"

func TestGetStatus(t *testing.T) {

	client := NewClient()

	err := client.GetStatus()

	if err != nil {
		t.Fatalf("There was an error getting status: %v", err)
	}
}

func TestGetMatches(t *testing.T) {

	client := NewClient()
	err := client.GetMatches()

	if err != nil {
		t.Fatalf("There was an error getting matches: %v", err)
	}
}

func TestGetMatch(t *testing.T) {
	client := NewClient()

	match, err := client.GetMatch("f78917d2-d7cf-11e6-ad79-062445d3d668")
	if err != nil {
		t.Fatalf("There was an error getting your match %v", err)
	}

	if match.ID != "f78917d2-d7cf-11e6-ad79-062445d3d668" {
		t.Fatalf("Got the wrong match")
	}
}

func TestGetPlayer(t *testing.T) {
	client := NewClient()

	player, err := client.GetPlayer("3a8f6440-1b20-11e6-a24a-06ab1a16f8e5")
	if err != nil {
		t.Fatalf("There was an error getting your player %v", err)
	}

	if player.ID != "3a8f6440-1b20-11e6-a24a-06ab1a16f8e5" {
		t.Fatalf("Got the wrong player")
	}
}
