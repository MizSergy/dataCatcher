package models

import (
	"github.com/kshvakov/clickhouse/lib/types"
	"time"
)

type FullTraffic struct {
	CreateAt        time.Time  `json:"create_at" db:"create_at,omitempty"`
	LeadCreate      time.Time  `json:"lead_create" db:"lead_create,omitempty"`
	CreateDate      time.Time  `json:"create_date" db:"create_date,omitempty"`
	VCode           string     `json:"vcode" db:"vcode,omitempty"`
	IsClick         int32      `json:"is_click" db:"is_click,omitempty"`
	SourceID        int32      `json:"source_id" db:"source_id,omitempty"`
	Campaign        int32      `json:"campaign" db:"campaign,omitempty"`
	StreamId        int32      `json:"stream_id" db:"stream_id,omitempty"`
	AffiliateID     int32      `json:"affiliate_id" db:"affiliate_id,omitempty"`
	PrelandID       int32      `json:"preland_id" db:"preland_id,omitempty"`
	IsBreaked       uint8      `json:"is_breaked" db:"is_breaked"`
	IsRefused       uint8      `json:"is_refused" db:"is_refused"`
	IsUnique        uint8      `json:"is_unique" db:"is_unique"`
	IsTest          uint8      `json:"is_test" db:"is_test"`
	ProcessInterval float64    `json:"process_interval" db:"process_interval"`
	ScreenHeight    int32      `json:"screen_height" db:"screen_height"`
	ScreenWidth     int32      `json:"screen_width" db:"screen_width"`
	Language        string     `json:"language" db:"language"`
	ClickPrice      float32    `json:"click_price" db:"click_price"`
	Browser         string     `json:"browser" db:"browser"`
	BrowserV        string     `json:"browserv" db:"browserv"`
	Os              string     `json:"os" db:"os"`
	OsV             string     `json:"osv" db:"osv"`
	Country         string     `json:"country" db:"country"`
	CountryCode     string     `json:"country_code" db:"country_code"`
	Region          string     `json:"region" db:"region"`
	City            string     `json:"city" db:"city"`
	Ip              uint32     `json:"ip" db:"ip"`
	Device          int32      `json:"device" db:"device"`
	IsMobil         uint8      `json:"is_mobil" db:"is_mobil"`
	Ad              string     `json:"ad" db:"ad"`
	Site            string     `json:"site" db:"site"`
	Sid1            string     `json:"sid1" db:"sid1"`
	Sid2            string     `json:"sid2" db:"sid2"`
	Sid3            string     `json:"sid3" db:"sid3"`
	Sid4            string     `json:"sid4" db:"sid4"`
	Sid5            string     `json:"sid5" db:"sid5"`
	Sid6            string     `json:"sid6" db:"sid6"`
	Sid7            string     `json:"sid7" db:"sid7"`
	PrelandUrl      string     `json:"preland_url" db:"preland_url"`
	Session         string     `json:"session" db:"session_id"`
	Url             string     `json:"url" db:"url"`
	Method          string     `json:"method" db:"method"`
	Params          string     `json:"params" db:"params"`
	StatusConfirmed int32      `json:"status_confirmed" db:"status_confirmed"`
	StatusHold      int32      `json:"status_hold" db:"status_hold"`
	StatusDeclined  int32      `json:"status_declined" db:"status_declined"`
	StatusOther     int32      `json:"status_other" db:"status_other"`
	StatusPaid      int32      `json:"status_paid" db:"status_paid"`
	OrderID         string     `json:"order_id" db:"order_id"`
	Amount          float32    `json:"amount" db:"amount"`
	ResultMessage   string     `json:"result_message" db:"result_message"`
	Profit          float32    `json:"profit" db:"profit"`
	PredictProfit   float32    `json:"predict_profit" db:"predict_profit"`
	AppearDate      types.Date `json:"appear_date" db:"appear_date"`
	Version         int8       `json:"version" db:"version"`
}
