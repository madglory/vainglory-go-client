package vainglory

import "time"

//Match structure
type Match struct {
	ID           string                 `json:"id" jsonapi:"primary,match" validate:"required"`
	CreatedAt    time.Time              `json:"createdAt" jsonapi:"attr,createdAt,iso8601" validate:"required"`
	Duration     int64                  `json:"duration" jsonapi:"attr,duration"`
	Rosters      []*Roster              `json:"rosters" jsonapi:"relation,rosters" validate:"required"`
	Rounds       []*Round               `json:"rounds" jsonapi:"relation,rounds"` // Rounds are optional @BJC
	Stats        map[string]interface{} `json:"stats" jsonapi:"attr,stats" validate:"required"`
	GameMode     string                 `json:"gameMode" jsonapi:"attr,gameMode"`
	PatchVersion string                 `json:"patchVersion" jsonapi:"attr,patchVersion"`
	TitleID      string                 `json:"titleId" jsonapi:"attr,titleId" validate:"required"`
	ShardID      string                 `json:"shardId" jsonapi:"attr,shardId" validate:"required"`
}
