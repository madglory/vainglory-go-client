package vainglory

import "time"

//Player structure
type Player struct {
	ID        string                 `json:"id" jsonapi:"primary,player" validate:"required"`
	Name      string                 `json:"name" jsonapi:"attr,name" validate:"required"`
	Stats     map[string]interface{} `json:"stats" jsonapi:"attr,stats"`
	TitleID   string                 `json:"titleId" jsonapi:"attr,titleId" validate:"required"`
	ShardID   string                 `json:"shardId" jsonapi:"attr,shardId" validate:"required"`
	CreatedAt time.Time              `json:"createdAt" jsonapi:"attr,createdAt,iso8601" validate:"required"`
}
