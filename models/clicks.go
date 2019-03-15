package models

import "time"

type Click struct {
	CreateAt time.Time `json:"create_at" db:"create_at"`
	VCode    string    `json:"vcode" db:"vcode"`
}
