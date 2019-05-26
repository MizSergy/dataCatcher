package controllers

import (
	"fmt"
	"redisCatcher/db"
	"redisCatcher/models"
	"time"
)

var breakData = make(map[string]*models.Breaking)
var pbData = make(map[string]map[string]models.PostBack)

func FillTraffic() {
	//fillClicks()
	//fillBreaks()
	fillLeads()
}

func fillClicks() {
	//index := 0
	//index2 := 1000000
	//
	//items := 1
	//for items > 0 {
	select_query := fmt.Sprintf(`SELECT 
	vcode,
	create_at,
 	source_id,
	campaign,
    preland_id,
    is_unique,
    is_test,
	br.stream_id,
	br.is_breaked,
    is_unique,
 	is_test,
	br.affiliate_id,
	br.process_interval,
	br.screen_width,
	br.screen_height,
	br.language,
	br.is_refused,
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
	br.create_date
FROM tracker_db.click_logs
ALL LEFT JOIN tracker_db.breaks as br FINAL USING vcode
PREWHERE toDate(create_at) BETWEEN '2019-03-05' and '2019-03-15'`)

	clickhouse := database.SqlxConnect()
	var collected_data []models.FullTraffic
	if err := clickhouse.Select(&collected_data, select_query); err != nil {
		fmt.Println(err)
	}
	clickhouse.Close()
	fmt.Println("Взяли")
	//items = len(collected_data)
	if len(collected_data) > 0 {
		time.Sleep(time.Second)
		//------------------------------------------Получаем клики из таблицы трафика-----------------------------------
		query :=
			`INSERT INTO tracker_db.traffic_data1
			( vcode,
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

		for _, item := range collected_data {
			if item.CreateAt.IsZero() {
				item.CreateAt = time.Now()
			}
			if item.CreateDate.IsZero() {
				item.CreateDate = time.Now()
			}
			if item.LeadCreate.IsZero() {
				item.LeadCreate = time.Now()
			}
			if _, err := stmt.Exec(
				item.VCode,
				item.CreateAt,
				item.CreateDate,
				item.LeadCreate,
				1,
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
	//fmt.Println("Чпоньк")
	//	index += index2 + 1
	//}
	fmt.Println("Запись кликов закончена")
}

func fillLeads() {
	index := 0
	index2 := 10000
	items := 1
	for items > 0 {

		pbData = make(map[string]map[string]models.PostBack)

		select_query := fmt.Sprintf(`SELECT * FROM tracker_db.post_backs PREWHERE toDate(create_at) BETWEEN '2019-03-02' AND '2019-04-01' AND LENGTH (vcode) = 36 ORDER BY create_at asc LIMIT %d,%d`, index, index2)
		var collected_data []models.PostBack
		var vcodeArray []string
		clickhouse := database.SqlxConnect()
		if err := clickhouse.Select(&collected_data, select_query); err != nil {
			fmt.Println(err)
		}
		clickhouse.Close()
		items = len(collected_data)
		index += index2 + 1

		for i := range collected_data {
			vcode := collected_data[i].VCode
			orderId := collected_data[i].OrderID
			_, ok := pbData[vcode]
			if !ok {
				vcodeArray = append(vcodeArray, collected_data[i].VCode)
				pbData[vcode] = make(map[string]models.PostBack)
				if len(orderId) == 0 {
					pbData[vcode][""] = collected_data[i]
					continue
				}
				pbData[vcode][orderId] = collected_data[i]
				continue
			}

			_, ok = pbData[vcode][orderId]
			if !ok {
				if len(orderId) == 0 {
					pbData[vcode][""] = collected_data[i]
					continue
				}
				pbData[vcode][orderId] = collected_data[i]
				continue
			}

			if collected_data[i].CreateAt.Sub(pbData[vcode][orderId].CreateAt) < 0 {
				continue
			}

			pbData[vcode][orderId] = collected_data[i]
		}

		var newTrafficArray []models.FullTraffic

		if len(vcodeArray) > 0 {
			//------------------------------------------Получаем клики из таблицы трафика-------------------------------
			trafficArray := GetTrafficData(database.SqlxConnect(), vcodeArray)
			if len(trafficArray) > 0 {
				oldTraffic := make([]models.FullTraffic, len(trafficArray))
				copy(oldTraffic, trafficArray)
				//------------------------------------------Мерджим данные--------------------------------------------------
				for i := range trafficArray {
					vcode := trafficArray[i].VCode
					orderId := trafficArray[i].OrderID
					_, ok := pbData[vcode]
					if !ok {
						continue
					}
					_, ok = pbData[vcode][orderId]
					if !ok {
						if len(orderId)==0{
							for _,v := range pbData[vcode]{
								if v.CreateAt.Sub(trafficArray[i].CreateAt) < 0 {
									continue
								}
								trafficArray[i] = v.TraffMerge(trafficArray[i])
								delete(pbData[vcode], v.OrderID)
								break
							}
						}
						continue
					}

					if trafficArray[i].CreateAt.Sub(pbData[vcode][orderId].CreateAt) < 0 {
						continue
					}

					trafficArray[i] = pbData[vcode][orderId].TraffMerge(trafficArray[i])
					delete(pbData[vcode], orderId)

					//if len(orderId) == 0 {
					//	trafficArray[i] = pbData[vcode][orderId].TraffMerge(trafficArray[i])
					//	delete(pbData[vcode], orderId)
					//	continue
					//} else {
					//
					//}
					//delete(pbData, vcode)
				}
				if len(oldTraffic) > 0 {
					RewriteTrafficData(oldTraffic, trafficArray)
				}
			}

			//-------------------------------------------Перегоняем новые(без дублей) данные в массив трафика---------------------------
			time.Sleep(time.Second)

			for _, val := range pbData {
				for _, item := range val {
					var newTraffic models.FullTraffic
					newTraffic.IsClick = 1
					newTrafficArray = append(newTrafficArray, item.TraffMerge(newTraffic))
					delete(pbData[item.VCode], item.OrderID)
				}
			}

			if len(newTrafficArray) > 0 {
				WriteTrafficData(newTrafficArray)
			}
		}
		fmt.Println("Чпоньк")
	}
	fmt.Println("Лиды ок")
}
