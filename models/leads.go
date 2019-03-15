package models

import "time"

type PostBack struct {
	VCode    			string			`json:"vcode" db:"vcode"`
	CreateAt 			time.Time		`json:"create_at" db:"create_at"`
	CreateDate 			time.Time		`json:"create_date" db:"create_date"`

	Url    				string			`json:"url" db:"url"`
	Method 				string			`json:"method" db:"method"`
	Params 				string			`json:"params" db:"params"`

	StatusConfirmed 	int32			`json:"status_confirmed" db:"status_confirmed"`
	StatusHold      	int32			`json:"status_hold" db:"status_hold"`
	StatusDeclined  	int32			`json:"status_declined" db:"status_declined"`
	StatusOther     	int32			`json:"status_other" db:"status_other"`

	OrderID 			string			`json:"order_id" db:"order_id"`
	Amount  			float32			`json:"amount" db:"amount"`

	Profit				float32			`json:"profit" db:"profit"`
	ResultMessage 		string			`json:"result_message" db:"result_message"`
}