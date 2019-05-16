package models

import (
	"time"
)

type BlackLists struct {
	Site      string    `form:"site,omitempty" json:"site,omitempty" bson:"site" db:"site"`
	SourceID  int       `form:"source_id,omitempty" json:"source_id,omitempty" bson:"source_id" db:"source_id"`
	CompanyID int       `form:"company_id,omitempty" json:"company_id,omitempty" bson:"company_id" db:"campaign"`
	Ban       uint8     `form:"ban" json:"ban" bson:"ban" db:"ban"`
	Version   int8      `db:"version"`
	CreateAt  time.Time `db:"create_at"`
}
