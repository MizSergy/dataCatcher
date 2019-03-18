package models

import "time"

type Click struct {
	CreateAt time.Time `json:"create_at" db:"create_at"`
	VCode    string    `json:"vcode" db:"vcode"`

	IsUnique uint8 `json:"is_unique" db:"is_unique"`
	IsMobil  uint8 `json:"is_mobil" db:"is_mobil"`
	Device   int32 `json:"device" db:"device"`

	Campaign   int32   `json:"campaign" db:"campaign"`
	SourceID   int32   `json:"source_id" db:"source_id"`
	ClickPrice float32 `json:"click_price" db:"click_price"`
	Browser    string  `json:"browser" db:"browser"`
	Os         string  `json:"os" db:"os"`
	Country    string  `json:"country" db:"country"`
	Region     string  `json:"region" db:"region"`
	City       string  `json:"city" db:"city"`
	Ip         uint32  `json:"ip" db:"ip"`
	Ad         string  `json:"ad" db:"ad"`
	Site       string  `json:"site" db:"site"`
	Sid1       string  `json:"sid1" db:"sid1"`
	Sid2       string  `json:"sid2" db:"sid2"`
	Sid3       string  `json:"sid3" db:"sid3"`
	Sid4       string  `json:"sid4" db:"sid4"`
	Sid5       string  `json:"sid5" db:"sid5"`
	PrelandUrl string  `json:"preland_url" db:"preland_url"`
	PrelandID  int32   `json:"preland_id" db:"preland_id"`
	Session    string  `json:"session" db:"session_id"`
	IsTest     uint8   `json:"is_test" db:"is_test"`

	CountryCode string `json:"country_code" db:"country_code"`
}
