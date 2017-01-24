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
