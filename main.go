package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"redisCatcher/controllers"
	"redisCatcher/db"
	"redisCatcher/logging"
	"redisCatcher/models"
)
//
//var leads = make(map[string]models.PostBack)
//var clicks = make(map[string]models.Click)
//var breaks = make(map[string]models.Breaks)
//var breakings = make(map[string]models.Breaks)
//
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ClickhouseInit()
	controllers.FillClicks()
	//CollectClicks()
	//controllers.Ban()
	//controllers.FillTraffic()
	//database.MongoConnect()
	//DataMerger()
}

func CollectClicks() {
	select_query := fmt.Sprintf(`
SELECT 
	*
FROM tracker_db.click_logs
PREWHERE toDate(create_at) BETWEEN '2019-03-05' and '2019-03-15'`)
	clickhouse := database.SqlxConnect()
	var collected_data []models.Click
	if err := clickhouse.Select(&collected_data, select_query); err != nil {
		fmt.Println(err)
	}
	clickhouse.Close()
	fmt.Println("Взяли")
	logging.Rewrite(collected_data)
}

//
//func DataMerger() {
//	//leadsKeys := getLeads()
//	getClicks()
//	//getBreakings(leadsKeys)
//	//getBreaks(click_keys)
//	//collectData()
//}
//
//func rewriteClick() {
//	var breaks_array []models.Breaks
//	for _, v := range clicks {
//		if brs, ok := breaks[v.VCode]; ok {
//			if v.CreateAt != brs.CreateAt && brs.IsBreaked != 1 {
//				item := models.Breaks{}
//				item = brs
//				stream := findFromMongo(v.Campaign,v.SourceID)
//				if brs.StreamId == 0{
//					item.StreamId = int32(stream[len(stream)-1].ID)
//				}
//				if brs.AffiliateID == 0{
//					item.AffiliateID = int32(stream[len(stream)-1].AffiliateID)
//				}
//				item.CreateAt = v.CreateAt
//				breaks_array = append(breaks_array, item)
//			}
//		} else {
//			item := models.Breaks{}
//			item.CreateAt = v.CreateAt
//			item.VCode = v.VCode
//			item.IsRefused = 0
//			item.IsBreaked = 0
//			stream := findFromMongo(v.Campaign,v.SourceID)
//			item.StreamId = int32(stream[len(stream)-1].ID)
//			item.AffiliateID = int32(stream[len(stream)-1].AffiliateID)
//			item.Version = 1
//			breaks_array = append(breaks_array, item)
//		}
//	}
//	if len(breaks_array) > 0 {
//		fmt.Println("Запись пробив")
//		query := `
//					INSERT INTO tracker_db.breaks
//						(vcode,
//						process_interval,
//						create_at,
//						stream_id,
//						is_breaked,
//						affiliate_id,
//						screen_width,
//						screen_height,
//						language,
//						is_refused,
//						version)
//					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
//
//		clickhouse := database.SqlxConnect()
//
//		tx, err := clickhouse.Begin()
//		ErrorCheck(err)
//
//		stmt, err := tx.Prepare(query)
//		ErrorCheck(err)
//		for _, item := range breaks_array {
//			if _, err := stmt.Exec(
//				item.VCode,
//				item.ProcessInterval,
//				item.CreateAt,
//				item.StreamId,
//				item.IsBreaked,
//				item.AffiliateID,
//				item.ScreenHeight,
//				item.ScreenWidth,
//				item.Language,
//				item.IsRefused,
//				1,
//			); err != nil {
//				fmt.Println(err.Error())
//			}
//
//		}
//		if err := tx.Commit(); err != nil {
//			fmt.Println(err.Error())
//		}
//		stmt.Close()
//		clickhouse.Close()
//	}
//	fmt.Println("Всё ок")
//	time.Sleep(time.Second * 2)
//}
//
////func collectData() {
////	breaks_array := []models.Breaks{}
////	for _, v := range leads {
////		if brs, ok := breaks[v.VCode]; ok {
////			if brs.AffiliateID == 0 || brs.StreamId != 2 {
////				var item models.Breaks
////				item = brs
////				stream_id, affiliate_id := findClick(brs.VCode)
////				item.IsBreaked = 1
////				item.AffiliateID = affiliate_id
////				item.StreamId = stream_id
////				item.Version = 1
////				breaks_array = append(breaks_array, item)
////			}
////		} else if brks, ok := breakings[v.VCode]; ok {
////			if brks.AffiliateID == 0 || brs.StreamId != 2 {
////				var item models.Breaks
////				item = brks
////				stream_id, affiliate_id := findClick(brks.VCode)
////				item.AffiliateID = affiliate_id
////				item.StreamId = stream_id
////				item.Version = 1
////				breaks_array = append(breaks_array, item)
////			}
////		} else if cls, ok := clicks[v.VCode]; ok {
////			var item models.Breaks
////			stream := findFromMongo(cls.Campaign, cls.SourceID)
////			item.VCode = cls.VCode
////			item.CreateAt = cls.CreateAt
////			item.Create_date = types.Date(cls.CreateAt)
////			item.IsRefused = 0
////			item.IsBreaked = 1
////			item.AffiliateID = int32(stream[len(stream)-1].AffiliateID)
////			item.StreamId = int32(stream[len(stream)-1 ].ID)
////			item.Version = 1
////			breaks_array = append(breaks_array, item)
////		} else {
////			fmt.Println("Клик с vcode=", v.VCode, " не найден")
////		}
////	}
////	fmt.Println(breaks_array)
////	if len(breaks_array) > 0 {
////		fmt.Println("Запись пробива")
////		query := `
////					INSERT INTO tracker_db.breaks
////						(vcode,
////						process_interval,
////						create_at,
////						stream_id,
////						is_breaked,
////						affiliate_id,
////						screen_width,
////						screen_height,
////						language,
////						is_refused,
////						version)
////					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
////
////		clickhouse := database.SqlxConnect()
////
////		tx, err := clickhouse.Begin()
////		ErrorCheck(err)
////
////		stmt, err := tx.Prepare(query)
////		ErrorCheck(err)
////		for _, item := range breaks_array {
////			if _, err := stmt.Exec(
////				item.VCode,
////				item.ProcessInterval,
////				item.CreateAt,
////				item.StreamId,
////				item.IsBreaked,
////				item.AffiliateID,
////				item.ScreenHeight,
////				item.ScreenWidth,
////				item.Language,
////				item.IsRefused,
////				1,
////			); err != nil {
////				fmt.Println(err.Error())
////			}
////
////		}
////		if err := tx.Commit(); err != nil {
////			fmt.Println(err.Error())
////		}
////
////		stmt.Close()
////		clickhouse.Close()
////	}
////	fmt.Println("Всё записали")
////
////}
//
//func findClick(vcode string) (stream_id, affiliate_id int32) {
//	if cls, ok := clicks[vcode]; ok {
//		item := findFromMongo(cls.Campaign, cls.SourceID)
//		return int32(item[len(item)-1].ID), int32(item[len(item)-1].AffiliateID)
//	} else {
//		return 0, 0
//	}
//
//}
//
//func findFromMongo(company_id int32, source_id int32) []models.Stream {
//	collection := database.MongoDB().C("stream")
//	//var result []company
//
//	result := []models.Stream{}
//	// Limit(100).
//	iter := collection.Find(bson.D{{"company_id", company_id}, {"source_id", 2}, {"type", 2}, {"is_archiving", false}}).Iter()
//	err := iter.All(&result)
//
//	if err = iter.Close(); err != nil {
//		panic(err)
//	}
//	return result
//}
//
//func getBreaks(keys []string) {
//	from := 0
//	to := 3000
//	max := len(keys)
//	for {
//		if to > max {
//			to = max
//		}
//		if from == max {
//			break
//		}
//		vcodeArray := keys[from:to]
//		vcodeString := ""
//		if len(vcodeArray) > 1 {
//			vcodeString = "'" + strings.Join(vcodeArray, "','") + "'"
//		} else {
//			vcodeString = "'" + vcodeArray[0] + "'"
//		}
//		query := fmt.Sprintf(`
//		SELECT * FROM tracker_db.breaks FINAL WHERE vcode in (%s)`, vcodeString)
//
//		clickhouse := database.SqlxConnect()
//		var collected_data []models.Breaks
//		if err := clickhouse.Select(&collected_data, query); err != nil {
//			fmt.Println(err)
//			break
//		}
//		time.Sleep(time.Second)
//		if len(collected_data) == 0 {
//			break
//		}
//
//		for _, v := range collected_data {
//			breaks[v.VCode] = v
//		}
//		from = to
//		to = to + 3000
//	}
//	fmt.Println("Получили breaks")
//}
//
////func getBreakings(keys []string) {
////	from := 0
////	to := 3000
////	max := len(keys)
////	for {
////		if to > max {
////			to = max
////		}
////		if from == max {
////			break
////		}
////		vcodeArray := keys[from:to]
////		vcodeString := ""
////		if len(vcodeArray) > 1 {
////			vcodeString = "'" + strings.Join(vcodeArray, "','") + "'"
////		} else {
////			vcodeString = "'" + vcodeArray[0] + "'"
////		}
////		query := fmt.Sprintf(`
////		SELECT * FROM tracker_db.breakings WHERE vcode in (%s)`, vcodeString)
////
////		clickhouse := database.SqlxConnect()
////		var collected_data []models.Breaks
////		if err := clickhouse.Select(&collected_data, query); err != nil {
////			fmt.Println(err)
////			break
////		}
////		time.Sleep(time.Second)
////		if len(collected_data) == 0 {
////			break
////		}
////
////		for _, v := range collected_data {
////			breakings[v.VCode] = v
////		}
////		from = to
////		to = to + 3000
////	}
////	fmt.Println("Получили breakings")
////}
//
//func getClicks(){
//		for {
//
//			query := `
//SELECT
//	vcode,
//	ld.create_date as create_date,
//	ld.url as url,
//	ld.method as method,
//	ld.params as params,
//	ld.status_confirmed as status_confirmed,
//	ld.status_hold as status_hold,
//	ld.status_declined as status_declined,
//	ld.status_other as status_other,
//	ld.order_id as order_id,
//	ld.amount as amount,
//	ld.result_message as result_message,
//	ld.version as version,
//	ld.profit as profit,
//	cl.create_at as create,
//	ld.create_at as ldcreate
//FROM tracker_db.leads as ld FINAL
//ALL LEFT JOIN tracker_db.click_logs as cl USING vcode
//WHERE toDate(cl.create_at) != toDate(ld.create_at) AND cl.is_test = 0 AND vcode != '' AND char_length(vcode)  > 20`
//
//			clickhouse := database.SqlxConnect()
//			var collected_data []models.PostBack
//			if err := clickhouse.Select(&collected_data, query); err != nil {
//				fmt.Println(err)
//			}
//			time.Sleep(time.Second)
//			if collected_data == nil || len(collected_data) == 0{
//				break
//			}
//			var breaks_array []models.PostBack
//
//			for _, v := range collected_data {
//				fmt.Println(v.Create,"!=",v.LDCreateDate)
//				if v.Create != v.LDCreateDate{
//					if v.VCode != ""{
//						item := models.PostBack{}
//						item.VCode = v.VCode
//						item.Url = v.Url
//						item.Method = v.Method
//						item.Params = v.Params
//						item.StatusConfirmed = v.StatusConfirmed
//						item.StatusDeclined = v.StatusDeclined
//						item.StatusHold = v.StatusHold
//						item.StatusOther = v.StatusOther
//						item.OrderID = v.OrderID
//						item.Amount = v.Amount
//						item.ResultMessage = v.ResultMessage
//						item.Version = 1
//						item.Profit = v.Profit
//						item.CreateDate = v.CreateDate
//						item.CreateAt = v.Create
//						fmt.Println(v.Create,"=",item.CreateAt)
//						breaks_array = append(breaks_array, item)
//					}
//				}
//
//
//			}
//			if len(breaks_array) > 0 {
//				fmt.Println("Запись пробив ", len(breaks_array))
//				query := `
//				INSERT INTO tracker_db.leads
//					(vcode,
//					create_at,
//					create_date,
//					url,
//					method,
//					params,
//					status_confirmed,
//					status_hold,
//					status_declined,
//					status_other,
//					order_id,
//					amount,
//					profit,
//					result_message,
//					version)
//				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
//
//				clickhouse := database.SqlxConnect()
//
//				tx, err := clickhouse.Begin()
//				ErrorCheck(err)
//
//				stmt, err := tx.Prepare(query)
//				ErrorCheck(err)
//				for _, item := range breaks_array {
//					if _, err := stmt.Exec(
//						item.VCode,
//						item.CreateAt,
//						item.CreateDate,
//						item.Url,
//						item.Method,
//						item.Params,
//						item.StatusConfirmed,
//						item.StatusHold,
//						item.StatusDeclined,
//						item.StatusOther,
//						item.OrderID,
//						item.Amount,
//						item.Profit,
//						item.ResultMessage,
//						1,
//					); err != nil {
//						fmt.Println(err.Error())
//					}
//
//				}
//				if err := tx.Commit(); err != nil {
//					fmt.Println(err.Error())
//				}
//				stmt.Close()
//				clickhouse.Close()
//			}
//			time.Sleep(time.Second *5)
//
//			//getBreaks(click_keys)
//			//rewriteClick()
//		}
//	fmt.Println("Всё закончилось")
//}
//
//func getLeads() []string {
//	var leads_keys []string
//	query := `
//		SELECT * FROM tracker_db.leads final WHERE vcode != ''
//		`
//	clickhouse := database.SqlxConnect()
//	var collected_data []models.PostBack
//	if err := clickhouse.Select(&collected_data, query); err != nil {
//		fmt.Println(err)
//	}
//	time.Sleep(time.Second)
//	for _, v := range collected_data {
//		leads[v.VCode] = v
//		leads_keys = append(leads_keys, v.VCode)
//	}
//	fmt.Println("Получили leads")
//	return leads_keys
//}

