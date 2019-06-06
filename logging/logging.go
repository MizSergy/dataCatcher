package logging

import (
	"fmt"
	"redisCatcher/db"
	"redisCatcher/models"
	"strings"
	"time"
)

type Order struct {
	LeadCreate      time.Time
	Url             string
	Method          string
	Params          string
	StatusConfirmed int32
	StatusHold      int32
	StatusDeclined  int32
	StatusOther     int32
	StatusPaid      int32
	OrderID         string
	Amount          float32
	ResultMessage   string
	Profit          float32
	PredictProfit   float32
}

type Item struct {
	VCode string

	CreateAt      time.Time
	CreateDate    time.Time
	IsClick       int32
	SourceID      int32
	Campaign      int32
	StreamId      int32
	AffiliateID   int32
	PrelandID     int32
	IsBreaked     uint8
	IsRefused     uint8
	IsUnique      uint8
	ProcessInterl float64
	ScreenHeight  int32
	ScreenWidth   int32
	Language      string
	ClickPrice    float32
	Browser       string
	BrowserV      string
	Os            string
	OsV           string
	Country       string
	CountryCode   string
	Region        string
	City          string
	Ip            uint32
	Device        int32
	IsMobil       uint8
	Ad            string
	Site          string
	Sid1          string
	Sid2          string
	Sid3          string
	Sid4          string
	Sid5          string
	Sid6          string
	Sid7          string
	PrelandUrl    string
	Session       string

	Orders   map[string]Order
	LifeTime time.Time
}

func (i *Item) AddOrUpdateOrder(o Order) {
	if i.Orders == nil {
		i.Orders = map[string]Order{}
	}
	i.Orders[o.OrderID] = o
}

//Vcode
type Log struct {
	UpdateVcode []string
	Items map[string]Item
}

type IGetVCode interface {
	GetVCode() string
}

type Merger interface {
	UpdateClick(vcode *string)
}

func (l *Log) Check(i IGetVCode) {
	if _, ok := l.Items[i.GetVCode()]; !ok{
		l.UpdateVcode = append(l.UpdateVcode, )
	}
}

func (l *Log) AddOrUpdate(i Item) {
	if l.Items == nil {
		l.Items = map[string]Item{}
	}
	// делаем запрос в бд для обновления cписка
	if len(l.UpdateVcode) > 0 {

		items := getTrafficData(l.UpdateVcode)
		for _,v :=range items{
			l.Items[i.VCode].TrafficUpdate(v)
		}
		l.UpdateVcode = []string{}
	}

	item, ok := l.Items[i.VCode]
	if !ok {
		l.Items[i.VCode] = i
	}
	if i.CreateAt.Sub(item.CreateAt) >= 0 {
		item.mergeItem(i)
	}
}

//Сравниваем входящие данные
func (i *Item) mergeItem(val Item) {
	// Переприсваиваем клики/пробивы
	i.CreateDate = val.CreateAt
	if val.IsUnique != 0 {
		i.IsUnique = val.IsUnique
	}
	if val.IsClick != 0 {
		i.IsClick = 1
	}
	i.Campaign = CheckOnEmptyInt(val.Campaign)
	i.SourceID = CheckOnEmptyInt(val.SourceID)
	if val.ClickPrice > 0 {
		i.ClickPrice = val.ClickPrice
	}
	if val.IsMobil > 0 {
		i.IsMobil = val.IsMobil
	}
	if val.Device > 0 {
		i.Device = val.Device
	}
	i.Browser = CheckOnEmptyStr(val.Browser)
	i.BrowserV = CheckOnEmptyStr(val.BrowserV)
	i.Os = CheckOnEmptyStr(val.Os)
	i.OsV = CheckOnEmptyStr(val.OsV)
	i.CountryCode = CheckOnEmptyStr(val.CountryCode)
	i.Country = CheckOnEmptyStr(val.Country)
	i.Region = CheckOnEmptyStr(val.Region)
	i.City = CheckOnEmptyStr(val.City)
	if val.Ip > 0 {
		i.Ip = val.Ip
	}
	i.Ad = CheckOnEmptyStr(val.Ad)
	i.Site = CheckOnEmptyStr(val.Site)
	i.Sid1 = CheckOnEmptyStr(val.Sid1)
	i.Sid2 = CheckOnEmptyStr(val.Sid2)
	i.Sid3 = CheckOnEmptyStr(val.Sid3)
	i.Sid4 = CheckOnEmptyStr(val.Sid4)
	i.Sid5 = CheckOnEmptyStr(val.Sid5)
	i.Sid6 = CheckOnEmptyStr(val.Sid6)
	i.Sid7 = CheckOnEmptyStr(val.Sid7)
	i.PrelandUrl = CheckOnEmptyStr(val.PrelandUrl)
	i.PrelandID = CheckOnEmptyInt(val.PrelandID)
	i.Session = CheckOnEmptyStr(val.Session)
	if val.IsBreaked != 0 {
		i.StreamId = CheckOnEmptyInt(val.StreamId)
		i.AffiliateID = CheckOnEmptyInt(val.AffiliateID)
	} else if val.IsBreaked == 0 && i.StreamId != 0{
		i.StreamId = CheckOnEmptyInt(val.StreamId)
		i.AffiliateID = CheckOnEmptyInt(val.AffiliateID)
	}
	if val.ProcessInterl > 0 {
		i.ProcessInterl = val.ProcessInterl
	}
	if val.ProcessInterl < 14 && val.ProcessInterl != 0 {
		i.IsRefused = 1
	}
	i.ScreenWidth = CheckOnEmptyInt(val.ScreenWidth)
	i.ScreenHeight = CheckOnEmptyInt(val.ScreenHeight)
	i.Language = CheckOnEmptyStr(val.Language)

	//Записываем данные постбека
	i.AddOrUpdateOrder(val.Orders[val.VCode])
}

func CheckOnEmptyStr(s string) string {
	if len(s) > 0 {
		return s
	}
	return ""
}

func CheckOnEmptyInt(i int32) int32 {
	if i >= 0 {
		return i
	}
	return 0
}

var AllTable Log

func (o Order) TrafficUpdate(t models.FullTraffic) {
	o.OrderID = CheckOnEmptyStr(t.OrderID)
	o.LeadCreate = t.LeadCreate
	o.Url = CheckOnEmptyStr(t.Url)
	o.Method = CheckOnEmptyStr(t.Method)
	o.Params = CheckOnEmptyStr(t.Params)

	o.StatusConfirmed = t.StatusConfirmed
	o.StatusHold = t.StatusHold
	o.StatusDeclined = t.StatusDeclined

	if t.StatusPaid != 0 {
		o.StatusPaid = t.StatusPaid
		o.StatusConfirmed = 1
	}

	o.ResultMessage = CheckOnEmptyStr(t.ResultMessage)
	if t.Amount != 0{
		o.Amount = t.Amount
	}
	if t.Profit != 0 && t.StatusConfirmed == 1{
		o.Profit = t.Amount
	}
	if t.PredictProfit == 0 && t.StatusHold == 1 {
		o.PredictProfit = t.Amount
}}

func (i Item) TrafficUpdate(t models.FullTraffic) {
	i.CreateDate = t.CreateAt
	if t.IsUnique != 0 {
		i.IsUnique = t.IsUnique
	}
	if t.IsClick != 0 {
		i.IsClick = 1
	}
	i.Campaign = CheckOnEmptyInt(t.Campaign)
	i.SourceID = CheckOnEmptyInt(t.SourceID)
	if t.ClickPrice > 0 {
		i.ClickPrice = t.ClickPrice
	}
	if t.IsMobil > 0 {
		i.IsMobil = t.IsMobil
	}
	if t.Device > 0 {
		i.Device = t.Device
	}
	i.Browser = CheckOnEmptyStr(t.Browser)
	i.BrowserV = CheckOnEmptyStr(t.BrowserV)
	i.Os = CheckOnEmptyStr(t.Os)
	i.OsV = CheckOnEmptyStr(t.OsV)
	i.CountryCode = CheckOnEmptyStr(t.CountryCode)
	i.Country = CheckOnEmptyStr(t.Country)
	i.Region = CheckOnEmptyStr(t.Region)
	i.City = CheckOnEmptyStr(t.City)
	if t.Ip > 0 {
		i.Ip = t.Ip
	}
	i.Ad = CheckOnEmptyStr(t.Ad)
	i.Site = CheckOnEmptyStr(t.Site)
	i.Sid1 = CheckOnEmptyStr(t.Sid1)
	i.Sid2 = CheckOnEmptyStr(t.Sid2)
	i.Sid3 = CheckOnEmptyStr(t.Sid3)
	i.Sid4 = CheckOnEmptyStr(t.Sid4)
	i.Sid5 = CheckOnEmptyStr(t.Sid5)
	i.Sid6 = CheckOnEmptyStr(t.Sid6)
	i.Sid7 = CheckOnEmptyStr(t.Sid7)
	i.PrelandUrl = CheckOnEmptyStr(t.PrelandUrl)
	i.PrelandID = CheckOnEmptyInt(t.PrelandID)
	i.Session = CheckOnEmptyStr(t.Session)
	if t.IsBreaked != 0 {
		i.StreamId = CheckOnEmptyInt(t.StreamId)
		i.AffiliateID = CheckOnEmptyInt(t.AffiliateID)
	} else if t.IsBreaked == 0 && i.StreamId != 0{
		i.StreamId = CheckOnEmptyInt(t.StreamId)
		i.AffiliateID = CheckOnEmptyInt(t.AffiliateID)
	}
	if t.ProcessInterval > 0 {
		i.ProcessInterl = t.ProcessInterval
	}
	if t.ProcessInterval < 14 && t.ProcessInterval != 0 {
		i.IsRefused = 1
	}
	i.ScreenWidth = CheckOnEmptyInt(t.ScreenWidth)
	i.ScreenHeight = CheckOnEmptyInt(t.ScreenHeight)
	i.Language = CheckOnEmptyStr(t.Language)

	//Записываем данные постбека
	i.Orders[t.VCode].TrafficUpdate(t)
}






func getTrafficData(vcodeArray []string) []models.FullTraffic {
	var vcodeString string
	if len(vcodeArray) > 1 {
		vcodeString = "'" + strings.Join(vcodeArray, "','") + "'"
	} else {
		vcodeString = "'" + vcodeArray[0] + "'"
	}
	clickhouse := database.SqlxConnect()
	select_query := fmt.Sprintf(`SELECT * FROM tracker_db.traffic_data1 FINAL PREWHERE vcode IN (%s)`, vcodeString)
	var collected_data []models.FullTraffic
	if err := clickhouse.Select(&collected_data, select_query); err != nil {
		fmt.Println(err)
	}
	clickhouse.Close()
	return collected_data
}