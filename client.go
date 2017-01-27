package vainglory

import (
	"github.com/google/jsonapi"

	gentleman "gopkg.in/h2non/gentleman.v1"
	"gopkg.in/h2non/gentleman.v1/plugins/query"
)

type client struct {
	g       *gentleman.Client
	baseURL string
}

//NewClient function exported
func NewClient(apikey string, q *QueryRequest) *client {
	c := &client{
		g:       gentleman.New(),
		baseURL: "https://api.dc01.gamelockerapp.com",
	}
	c.g.BaseURL(c.baseURL)
	c.g.SetHeader("Client", "gentleman")
	c.g.SetHeader("Authorization", "Bearer "+apikey)
	c.g.SetHeader("X-TITLE-ID", "semc-vainglory")
	c.g.SetHeader("Accept", "application/vnd.api+json")
	c.g.Use(query.Set("page[limit]", q.Limit))
	c.g.Use(query.Set("page[offset]", q.Offset))
	c.g.Use(query.Set("sort", q.SortField))
	c.g.Use(query.Set("filter[createdAt-start]", q.CreatedAtStart))
	c.g.Use(query.Set("filter[createdAt-end]", q.CreatedAtEnd))
	c.g.Use(query.Set("filter[playerNames]", q.PlayerNames))
	c.g.Use(query.Set("filer[teamNames]", q.TeamNames))

	return c
}

func (c *client) GetStatus() (string, int, error) {

	req := c.g.Request()
	req.Path("/status")
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return "", 0, err
	}
	if !res.Ok {
		return "", res.StatusCode, err
	}

	return res.String(), 0, err
}

func (c *client) GetMatches() (string, int, error) {

	req := c.g.Request()
	req.Path("/shards/na/matches")

	req.Method("GET")
	res, err := req.Send()
	if err != nil {
		return "", 0, err
	}
	if !res.Ok {
		return "", res.StatusCode, err
	}

	return res.String(), 0, err
}

func (c *client) GetMatchByID(matchID string) (*Match, int, error) {
	req := c.g.Request()
	req.Path("/shards/na/matches/" + matchID)
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return nil, 0, err
	}

	if !res.Ok {
		return nil, res.StatusCode, err
	}

	match := new(Match)
	if err = jsonapi.UnmarshalPayload(res.RawResponse.Body, match); err != nil {
		return nil, 0, err
	}
	return match, 0, err
}

func (c *client) GetPlayerByID(playerID string) (*Player, int, error) {

	req := c.g.Request()
	req.Path("/shards/na/players/" + playerID)
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return nil, 0, err
	}

	if !res.Ok {
		return nil, res.StatusCode, err
	}

	player := new(Player)
	if err = jsonapi.UnmarshalPayload(res.RawResponse.Body, player); err != nil {
		return nil, 0, err
	}
	return player, 0, err
}
