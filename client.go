package vainglory

import (
	"fmt"

	"github.com/google/jsonapi"

	gentleman "gopkg.in/h2non/gentleman.v1"
)

type client struct {
	g       *gentleman.Client
	baseURL string
}

//NewClient function exported
func NewClient() *client {
	c := &client{
		g:       gentleman.New(),
		baseURL: "https://api.dc01.gamelockerapp.com",
	}
	c.g.BaseURL(c.baseURL)
	c.g.SetHeader("Client", "gentleman")
	c.g.SetHeader("Authorization", "Bearer aaa.bbb.ccc")
	c.g.SetHeader("X-TITLE-ID", "semc-vainglory")
	c.g.SetHeader("Accept", "application/vnd.api+json")

	return c
}

func (c *client) GetStatus() error {

	req := c.g.Request()
	req.Path("/status")
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return err
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return err
	}

	//fmt.Printf("Status: %v", res.String())
	return err
}

func (c *client) GetMatches() error {

	req := c.g.Request()
	req.Path("/shards/na/matches")
	req.Method("GET")
	//req.Param("page[limit]", "2")

	res, err := req.Send()
	if err != nil {
		return err
	}
	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return err
	}

	//fmt.Printf("Get matches: %s", res.String())
	return err
}

func (c *client) GetPlayer(playerID string) (*Player, error) {

	req := c.g.Request()
	req.Path("/shards/na/players/" + playerID)
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return nil, err
	}

	if !res.Ok {
		fmt.Printf("Invalid server response: %d\n", res.StatusCode)
		return nil, err
	}

	//create an empty player object
	player := new(Player)

	//fill the player object with the response data
	if err = jsonapi.UnmarshalPayload(res.RawResponse.Body, player); err != nil {
		fmt.Printf("Error %s", err)
		return nil, err
	}
	//fmt.Print(prettyPlayer(player))
	return player, err
}
