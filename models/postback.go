package models

import "C"
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

func (c PostBack) Merge(f FullTraffic) FullTraffic {
	if c.CreateDate.IsZero() {
		c.CreateDate = time.Now()
	}
	if c.CreateDate.Sub(f.CreateAt) < 0 && c.OrderID == f.OrderID {
		return f
	}
	if f.OrderID != c.OrderID {
		if (f.IsClick == 1 && len(f.OrderID) != 0) || len(f.VCode) == 0 {
			f.IsClick = 0
		} else {
			f.IsClick = 1
		}
	} else {
		f.IsClick = 1
	}

	f.OrderID = c.OrderID
	f.CreateAt = c.CreateAt
	f.VCode = c.VCode
	f.LeadCreate = c.CreateDate
	f.CreateDate = c.CreateDate
	f.Url = c.Url
	f.Method = c.Method
	f.Params = c.Params
	f.StatusConfirmed = c.StatusConfirmed
	f.StatusHold = c.StatusHold
	f.StatusDeclined = c.StatusDeclined
	f.StatusOther = c.StatusOther
	f.StatusPaid = c.StatusPaid

	if c.StatusConfirmed == 1 {
		f.Profit = c.Amount
	}

	f.Amount = c.Amount
	f.ResultMessage = c.ResultMessage
	if c.PredictProfit == 0 && c.StatusHold == 1 {
		f.PredictProfit = c.Amount
		return f
	}
	f.PredictProfit = c.PredictProfit
	return f
}