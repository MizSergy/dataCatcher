package models

import (
	"github.com/kshvakov/clickhouse/lib/types"
	"redisCatcher/logging"
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

func (b Breaking) Merge(f FullTraffic) FullTraffic {
	if b.IsBreaked != 0 {
		f.StreamId = logging.CheckOnEmptyInt(b.StreamId)
		f.AffiliateID = logging.CheckOnEmptyInt(b.AffiliateID)
	} else if b.IsBreaked == 0 && f.StreamId != 0{
		f.StreamId = logging.CheckOnEmptyInt(b.StreamId)
		f.AffiliateID = logging.CheckOnEmptyInt(b.AffiliateID)
	}
	if b.ProcessInterval > 0 {
		f.ProcessInterval = b.ProcessInterval
	}
	if b.ProcessInterval < 14 && b.ProcessInterval != 0 {
		f.IsRefused = 1
	}
	f.ScreenWidth = logging.CheckOnEmptyInt(b.ScreenWidth)
	f.ScreenHeight = logging.CheckOnEmptyInt(b.ScreenHeight)
	f.Language = logging.CheckOnEmptyStr(b.Language)
	return f
}