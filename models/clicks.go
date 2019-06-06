package models

import (
	"redisCatcher/logging"
	"time"
)

type Click struct {
	CreateAt time.Time `json:"create_at" db:"create_at"`
	VCode    string    `json:"vcode" db:"vcode"`
	IsUnique uint8 `json:"is_unique" db:"is_unique"`
	IsMobil  uint8 `json:"is_mobil" db:"is_mobil"`
	Device   int32 `json:"device" db:"device"`
	Campaign    int32   `json:"campaign" db:"campaign"`
	SourceID    int32   `json:"source_id" db:"source_id"`
	AffiliateID int32   `json:"affiliate_id" db:"affiliate_id"`
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

func (c Click) GetVCode() string {
	return c.VCode
}

func (c Click) Merge(f FullTraffic) FullTraffic {
	// Переприсваиваем клики
	if c.IsUnique != 0 {
		f.IsUnique = c.IsUnique
	}
	f.Campaign = logging.CheckOnEmptyInt(c.Campaign)
	f.SourceID = logging.CheckOnEmptyInt(c.SourceID)
	if c.ClickPrice > 0 {
		f.ClickPrice = c.ClickPrice
	}
	if c.IsMobil > 0 {
		f.IsMobil = c.IsMobil
	}
	if c.Device > 0 {
		f.Device = c.Device
	}
	f.Browser = logging.CheckOnEmptyStr(c.Browser)
	f.BrowserV = logging.CheckOnEmptyStr(c.BrowserV)
	f.Os = logging.CheckOnEmptyStr(c.Os)
	f.OsV = logging.CheckOnEmptyStr(c.OsV)
	f.CountryCode = logging.CheckOnEmptyStr(c.CountryCode)
	f.Country = logging.CheckOnEmptyStr(c.Country)
	f.Region = logging.CheckOnEmptyStr(c.Region)
	f.City = logging.CheckOnEmptyStr(c.City)
	if c.Ip > 0 {
		f.Ip = c.Ip
	}
	f.Ad = logging.CheckOnEmptyStr(c.Ad)
	f.Site = logging.CheckOnEmptyStr(c.Site)
	f.Sid1 = logging.CheckOnEmptyStr(c.Sid1)
	f.Sid2 = logging.CheckOnEmptyStr(c.Sid2)
	f.Sid3 = logging.CheckOnEmptyStr(c.Sid3)
	f.Sid4 = logging.CheckOnEmptyStr(c.Sid4)
	f.Sid5 = logging.CheckOnEmptyStr(c.Sid5)
	f.Sid6 = logging.CheckOnEmptyStr(c.Sid6)
	f.Sid7 = logging.CheckOnEmptyStr(c.Sid7)
	f.PrelandUrl = logging.CheckOnEmptyStr(c.PrelandUrl)
	f.PrelandID = logging.CheckOnEmptyInt(c.PrelandID)
	f.Session = logging.CheckOnEmptyStr(c.Session)
	return f
}

func (c Click) UpdateClick(vcode string) {
	vcode = c.VCode
}
