package vainglory

import (
	"io"
	"reflect"

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
	c.g.SetHeader("Content-Encoding", "gzip")

	c.g.Use(query.Set("page[limit]", q.Limit))
	c.g.Use(query.Set("page[offset]", q.Offset))
	c.g.Use(query.Set("sort", q.SortField))
	c.g.Use(query.Set("filter[createdAt-start]", q.CreatedAtStart))
	c.g.Use(query.Set("filter[createdAt-end]", q.CreatedAtEnd))
	c.g.Use(query.Set("filter[playerNames]", q.PlayerNames))
	c.g.Use(query.Set("filer[teamNames]", q.TeamNames))

	return c
}

func (c *client) getRequest(path string) (io.Reader, int, error) {

	req := c.g.Request()
	req.Path(path)
	req.Method("GET")

	res, err := req.Send()
	if err != nil {
		return nil, 0, err
	}

	if !res.Ok {
		return nil, res.StatusCode, err
	}

	return res.RawResponse.Body, 0, nil
}

func (c *client) GetStatus() (io.Reader, int, error) {

	res, serverResponse, error := c.getRequest("/status")

	return res, serverResponse, error
}

func (c *client) GetMatches() ([]*Match, int, error) {

	res, serverResponse, err := c.getRequest("/shards/na/matches")

	data, unmarshalErr := jsonapi.UnmarshalManyPayload(res, reflect.TypeOf(new(Match)))
	if unmarshalErr != nil {
		return nil, 0, err
	}
	matches := []*Match{}
	for _, b := range data {
		match, ok := b.(*Match)
		if !ok {
			return nil, 0, err
		}
		matches = append(matches, match)
	}
	return matches, serverResponse, err
}

func (c *client) GetMatchByID(matchID string) (*Match, int, error) {

	res, serverResponse, err := c.getRequest("/shards/na/matches/" + matchID)
	match := new(Match)
	if unmarshalErr := jsonapi.UnmarshalPayload(res, match); unmarshalErr != nil {
		return nil, 0, unmarshalErr
	}
	return match, serverResponse, err
}

func (c *client) GetPlayerByID(playerID string) (*Player, int, error) {

	res, serverResponse, err := c.getRequest("/shards/na/players/" + playerID)

	player := new(Player)
	if unmarshalErr := jsonapi.UnmarshalPayload(res, player); unmarshalErr != nil {
		return nil, 0, err
	}
	return player, serverResponse, err
}
