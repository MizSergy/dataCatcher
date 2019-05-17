package models

import (
	"time"
)

type PostBack struct {
	VCode      string    `json:"vcode" db:"vcode"`
	CreateAt   time.Time `json:"create_at" db:"create_at"`
	CreateDate time.Time `json:"create_date" db:"create_date"`

	Url    string `json:"url" db:"url"`
	Method string `json:"method" db:"method"`
	Params string `json:"params" db:"params"`

	StatusConfirmed int32 `json:"status_confirmed" db:"status_confirmed"`
	StatusHold      int32 `json:"status_hold" db:"status_hold"`
	StatusDeclined  int32 `json:"status_declined" db:"status_declined"`
	StatusOther     int32 `json:"status_other" db:"status_other"`
	StatusPaid      int32 `json:"status_paid" db:"status_paid"`

	OrderID string  `json:"order_id" db:"order_id"`
	Amount  float32 `json:"amount" db:"amount"`

	Profit        float32 `json:"profit" db:"profit"`
	PredictProfit float32 `json:"predict_profit" db:"predict_profit"`
	ResultMessage string  `json:"result_message" db:"result_message"`
}

func (c PostBack) TraffMerge(val FullTraffic) FullTraffic {
	if c.CreateDate.IsZero() {
		c.CreateDate = time.Now()
	}
	if c.CreateDate.Sub(val.CreateAt) >= 0 {
		if val.OrderID != c.OrderID && val.IsClick != 0{
			val.IsClick = 0
		} else {
			val.IsClick = 1
		}
		val.CreateAt = c.CreateAt
		val.VCode = c.VCode
		val.LeadCreate = c.CreateDate
		val.CreateDate = c.CreateDate
		val.Url = c.Url
		val.Method = c.Method
		val.Params = c.Params
		val.StatusConfirmed = c.StatusConfirmed
		val.StatusHold = c.StatusHold
		val.StatusDeclined = c.StatusDeclined
		val.StatusOther = c.StatusOther
		val.OrderID = c.OrderID
		val.Amount = c.Amount
		val.ResultMessage = c.ResultMessage
		val.PredictProfit = c.PredictProfit
	}
	return val
}
