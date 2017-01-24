package vainglory

//Round structure
type Round struct {
	ID       string                 `json:"id" jsonapi:"primary,round"  validate:"required"`
	Ordinal  int64                  `json:"ordinal" jsonapi:"attr,ordinal"  validate:"required"`
	Stats    map[string]interface{} `json:"stats" jsonapi:"attr,stats"  validate:"required"`
	Duration int64                  `json:"duration" jsonapi:"attr,duration"`
}
