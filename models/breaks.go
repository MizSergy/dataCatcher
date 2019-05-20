package models

import (
	"github.com/kshvakov/clickhouse/lib/types"
	"time"
)

type Breaking struct {
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

func (c Breaking) Merge(val FullTraffic) FullTraffic {
	if c.ProcessInterval > 0{
		val.ProcessInterval = c.ProcessInterval
	}
	if c.ScreenWidth != 0{
		val.ScreenWidth = c.ScreenWidth
	}
	if c.ScreenHeight != 0{
		val.ScreenHeight = c.ScreenHeight
	}
	if len(c.Language) != 0 {
		val.Language = c.Language
	}

	if c.IsBreaked != 0{
		val.IsBreaked = c.IsBreaked
		if c.CreateAt.Sub(val.CreateAt) >= 0 {
			val.CreateAt = c.CreateAt
			if c.StreamId != 0{
				val.StreamId = c.StreamId
			}
			if c.AffiliateID != 0{
				val.AffiliateID = c.AffiliateID
			}
		}
	} else {
		if c.StreamId != 0{
			val.StreamId = c.StreamId
		}
		if c.AffiliateID != 0{
			val.AffiliateID = c.AffiliateID
		}
	}

	if c.ProcessInterval != 0 && c.ProcessInterval < 14{
		val.IsRefused = 1
	}

	return val
}
