package controllers

import (
	"redisCatcher/db"
	"redisCatcher/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type DataMerger interface {
	Merge(traffic models.FullTraffic) models.FullTraffic
}

func GetTrafficData(clickhouse *sqlx.DB, vcodeArray []string) []models.FullTraffic {
	var vcodeString string
	if len(vcodeArray) > 1 {
		vcodeString = "'" + strings.Join(vcodeArray, "','") + "'"
	} else {
		vcodeString = "'" + vcodeArray[0] + "'"
	}

	select_query := fmt.Sprintf(`SELECT * FROM tracker_db.traffic_data1 FINAL PREWHERE vcode IN (%s)`, vcodeString)
	var collected_data []models.FullTraffic
	if err := clickhouse.Select(&collected_data, select_query); err != nil {
		fmt.Println(err)
	}
	clickhouse.Close()
	return collected_data
}

func RewriteTrafficData(oldData []models.FullTraffic, newData []models.FullTraffic){
	query :=
		`INSERT INTO tracker_db.traffic_data1
			(
 			 vcode,
 			 create_at,
		  	 create_date,
 	 		 lead_create,
			 is_click,
 	 		 source_id,
	  		 campaign,
 	 		 stream_id,
 	 		 affiliate_id,
 	 		 preland_id,
 	 		 is_breaked,
 	 		 is_refused,
	  		 is_unique,
 	 		 is_test,
 	 		 process_interval,
 	 		 screen_width,
 	 		 screen_height,
			 language,
		 	 click_price,
	  		 browser,
	  		 browserv,
			 os,
			 osv,
		  	 country,
			 country_code,
  			 region,
	  	 	 city,
 	 		 ip,
		  	 device,
  			 is_mobil,
 		 	 ad,
 		     site,
 			 sid1,
 			 sid2,
 			 sid3,
 		 	 sid4,
 		 	 sid5,
 		 	 sid6,
 		 	 sid7,
 		 	 preland_url,
		  	 session_id,
		  	 url,
 		 	 method,
  		 	 params,
  			 status_confirmed,
 			 status_hold,
  			 status_declined,
	  	 	 status_other,
			 status_paid,
 			 order_id,
 			 amount,
 			 result_message,
 			 predict_profit,
  			 profit,
 			 version)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	clickhouse_conn := database.SqlxConnect()

	tx, err := clickhouse_conn.Begin()
	ErrorCheck(err)

	stmt, err := tx.Prepare(query)
	ErrorCheck(err)

	for i, item := range newData {
		if _, err := stmt.Exec(
			oldData[i].VCode,
			oldData[i].CreateAt,
			oldData[i].CreateDate,
			oldData[i].LeadCreate,
			oldData[i].IsClick,
			oldData[i].SourceID,
			oldData[i].Campaign,
			oldData[i].StreamId,
			oldData[i].AffiliateID,
			oldData[i].PrelandID,
			oldData[i].IsBreaked,
			oldData[i].IsRefused,
			oldData[i].IsUnique,
			oldData[i].IsTest,
			oldData[i].ProcessInterval,
			oldData[i].ScreenWidth,
			oldData[i].ScreenHeight,
			oldData[i].Language,
			oldData[i].ClickPrice,
			oldData[i].Browser,
			oldData[i].BrowserV,
			oldData[i].Os,
			oldData[i].OsV,
			oldData[i].Country,
			oldData[i].CountryCode,
			oldData[i].Region,
			oldData[i].City,
			oldData[i].Ip,
			oldData[i].Device,
			oldData[i].IsMobil,
			oldData[i].Ad,
			oldData[i].Site,
			oldData[i].Sid1,
			oldData[i].Sid2,
			oldData[i].Sid3,
			oldData[i].Sid4,
			oldData[i].Sid5,
			oldData[i].Sid6,
			oldData[i].Sid7,
			oldData[i].PrelandUrl,
			oldData[i].Session,
			oldData[i].Url,
			oldData[i].Method,
			oldData[i].Params,
			oldData[i].StatusConfirmed,
			oldData[i].StatusHold,
			oldData[i].StatusDeclined,
			oldData[i].StatusOther,
			oldData[i].StatusPaid,
			oldData[i].OrderID,
			oldData[i].Amount,
			oldData[i].ResultMessage,
			oldData[i].PredictProfit,
			oldData[i].Profit,
			-1,
		); err != nil {
			fmt.Println(err.Error())
		}
		if _, err := stmt.Exec(
			item.VCode,
			item.CreateAt,
			item.CreateDate,
			item.LeadCreate,
			item.IsClick,
			item.SourceID,
			item.Campaign,
			item.StreamId,
			item.AffiliateID,
			item.PrelandID,
			item.IsBreaked,
			item.IsRefused,
			item.IsUnique,
			item.IsTest,
			item.ProcessInterval,
			item.ScreenWidth,
			item.ScreenHeight,
			item.Language,
			item.ClickPrice,
			item.Browser,
			item.BrowserV,
			item.Os,
			item.OsV,
			item.Country,
			item.CountryCode,
			item.Region,
			item.City,
			item.Ip,
			item.Device,
			item.IsMobil,
			item.Ad,
			item.Site,
			item.Sid1,
			item.Sid2,
			item.Sid3,
			item.Sid4,
			item.Sid5,
			item.Sid6,
			item.Sid7,
			item.PrelandUrl,
			item.Session,
			item.Url,
			item.Method,
			item.Params,
			item.StatusConfirmed,
			item.StatusHold,
			item.StatusDeclined,
			item.StatusOther,
			item.StatusPaid,
			item.OrderID,
			item.Amount,
			item.ResultMessage,
			item.PredictProfit,
			item.Profit,
			1,
		); err != nil {
			fmt.Println(err.Error())
		}
	}
	if err := tx.Commit(); err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	clickhouse_conn.Close()
}


func WriteTrafficData(newData []models.FullTraffic){
	query :=
		`INSERT INTO tracker_db.traffic_data1
		(	 vcode,
 			 create_at,
		  	 create_date,
 	 		 lead_create,
			 is_click,
 	 		 source_id,
	  		 campaign,
 	 		 stream_id,
 	 		 affiliate_id,
 	 		 preland_id,
 	 		 is_breaked,
 	 		 is_refused,
	  		 is_unique,
 	 		 is_test,
 	 		 process_interval,
 	 		 screen_width,
 	 		 screen_height,
			 language,
		 	 click_price,
	  		 browser,
	  		 browserv,
			 os,
			 osv,
		  	 country,
			 country_code,
  			 region,
	  	 	 city,
 	 		 ip,
		  	 device,
  			 is_mobil,
 		 	 ad,
 		     site,
 			 sid1,
 			 sid2,
 			 sid3,
 		 	 sid4,
 		 	 sid5,
 		 	 sid6,
 		 	 sid7,
 		 	 preland_url,
		  	 session_id,
		  	 url,
 		 	 method,
  		 	 params,
  			 status_confirmed,
 			 status_hold,
  			 status_declined,
	  	 	 status_other,
	  	 	 status_paid,
 			 order_id,
 			 amount,
 			 result_message,
 			 predict_profit,
  			 profit,
 			 version)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	clickhouse_conn := database.SqlxConnect()

	tx, err := clickhouse_conn.Begin()
	ErrorCheck(err)

	stmt, err := tx.Prepare(query)
	ErrorCheck(err)

	for _, item := range newData {
		if item.CreateAt.IsZero(){
			item.CreateAt = time.Now()
		}
		if item.CreateDate.IsZero(){
			item.CreateDate = time.Now()
		}
		if item.LeadCreate.IsZero(){
			item.LeadCreate = time.Now()
		}
		if _, err := stmt.Exec(
			item.VCode,
			item.CreateAt,
			item.CreateDate,
			item.LeadCreate,
			item.IsClick,
			item.SourceID,
			item.Campaign,
			item.StreamId,
			item.AffiliateID,
			item.PrelandID,
			item.IsBreaked,
			item.IsRefused,
			item.IsUnique,
			item.IsTest,
			item.ProcessInterval,
			item.ScreenWidth,
			item.ScreenHeight,
			item.Language,
			item.ClickPrice,
			item.Browser,
			item.BrowserV,
			item.Os,
			item.OsV,
			item.Country,
			item.CountryCode,
			item.Region,
			item.City,
			item.Ip,
			item.Device,
			item.IsMobil,
			item.Ad,
			item.Site,
			item.Sid1,
			item.Sid2,
			item.Sid3,
			item.Sid4,
			item.Sid5,
			item.Sid6,
			item.Sid7,
			item.PrelandUrl,
			item.Session,
			item.Url,
			item.Method,
			item.Params,
			item.StatusConfirmed,
			item.StatusHold,
			item.StatusDeclined,
			item.StatusOther,
			item.StatusPaid,
			item.OrderID,
			item.Amount,
			item.ResultMessage,
			item.PredictProfit,
			item.Profit,
			1,
		); err != nil {
			fmt.Println(err.Error())
		}
	}
	if err := tx.Commit(); err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	clickhouse_conn.Close()
}