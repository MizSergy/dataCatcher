package models

import "time"

type Click struct {
	CreateAt time.Time `json:"create_at" db:"create_at"`
	VCode    string    `json:"vcode" db:"vcode"`

	IsUnique uint8 `json:"is_unique" db:"is_unique"`
	IsMobil  uint8 `json:"is_mobil" db:"is_mobil"`
	Device   int32 `json:"device" db:"device"`

	Campaign    int32   `json:"campaign" db:"campaign"`
	SourceID    int32   `json:"source_id" db:"source_id"`
	ClickPrice  float32 `json:"click_price" db:"click_price"`
	Browser     string  `json:"browser" db:"browser"`
	BrowserV    string  `json:"browserv" db:"browserv"`
	Os          string  `json:"os" db:"os"`
	OsV         string  `json:"osv" db:"osv"`
	Country     string  `json:"country" db:"country"`
	Region      string  `json:"region" db:"region"`
	City        string  `json:"city" db:"city"`
	Ip          uint32  `json:"ip" db:"ip"`
	Ad          string  `json:"ad" db:"ad"`
	Site        string  `json:"site" db:"site"`
	Sid1        string  `json:"sid1" db:"sid1"`
	Sid2        string  `json:"sid2" db:"sid2"`
	Sid3        string  `json:"sid3" db:"sid3"`
	Sid4        string  `json:"sid4" db:"sid4"`
	Sid5        string  `json:"sid5" db:"sid5"`
	Sid6        string  `json:"sid6" db:"sid6"`
	Sid7        string  `json:"sid7" db:"sid7"`
	PrelandUrl  string  `json:"preland_url" db:"preland_url"`
	PrelandID   int32   `json:"preland_id" db:"preland_id"`
	Session     string  `json:"session" db:"session_id"`
	IsTest      uint8   `json:"is_test" db:"is_test"`
	CountryCode string  `json:"country_code" db:"country_code"`
}

func (c Click) Merge(val FullTraffic) FullTraffic {
	val.VCode = c.VCode
	val.CreateAt = c.CreateAt
	val.CreateAt = c.CreateAt
	val.CreateAt = c.CreateAt
	val.IsUnique = c.IsUnique
	val.IsTest = c.IsTest
	val.Campaign = c.Campaign
	val.SourceID = c.SourceID
	val.ClickPrice = c.ClickPrice
	val.IsMobil = c.IsMobil
	val.Device = c.Device
	val.Browser = c.Browser
	val.BrowserV = c.BrowserV
	val.Os = c.Os
	val.OsV = c.OsV
	val.CountryCode = c.CountryCode
	val.Country = c.Country
	val.Region = c.Region
	val.City = c.City
	val.Ip = c.Ip
	val.Ad = c.Ad
	val.Site = c.Site
	val.Sid1 = c.Sid1
	val.Sid2 = c.Sid2
	val.Sid3 = c.Sid3
	val.Sid4 = c.Sid4
	val.Sid5 = c.Sid5
	val.Sid6 = c.Sid6
	val.Sid7 = c.Sid7
	val.PrelandUrl = c.PrelandUrl
	val.PrelandID = c.PrelandID
	val.Session = c.Session
	return val
}