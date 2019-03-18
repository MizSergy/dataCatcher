package models

import (
	"github.com/kshvakov/clickhouse/lib/types"
	"time"
)

type Breaks struct {
	VCode    string    `json:"vcode" db:"vcode"`
	CreateAt time.Time `json:"create_at" db:"create_at"`

	IsBreaked       uint8   `json:"is_breaked" db:"is_breaked"`
	StreamId        int32   `json:"stream_id" db:"stream_id"`
	ProcessInterval float64 `json:"process_interval" db:"process_interval"`
	AffiliateID     int32   `json:"affiliate_id" db:"affiliate_id"`

	ScreenHeight int32      `json:"screen_height" db:"screen_height"`
	ScreenWidth  int32      `json:"screen_width" db:"screen_width"`
	Language     string     `json:"language" db:"language"`
	IsRefused    uint8      `json:"is_refused" db:"is_refused"`
	Version      int8       `json:"version" db:"version"`
	Create_date  types.Date `json:"create_date" db:"create_date"`

	Server string `json:"server_number"`
}