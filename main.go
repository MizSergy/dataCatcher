package main

import (
	"dataCatcher/db"
	"dataCatcher/models"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kshvakov/clickhouse/lib/types"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"strings"
	"time"
)

var leads = make(map[string]models.PostBack)
var clicks = make(map[string]models.Click)
var breaks = make(map[string]models.Breaks)
var breakings = make(map[string]models.Breaks)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ClickhouseInit()
	database.MongoConnect()

	DataMerger()
}

func DataMerger() {
	leadsKeys := getLeads()
	getClicks(leadsKeys)
	getBreakings(leadsKeys)
	getBreaks(leadsKeys)
	collectData()
}

func collectData() {
	breaks_array := []models.Breaks{}
	for _, v := range leads {
		if brs, ok := breaks[v.VCode]; ok {
			if brs.AffiliateID == 0 || brs.StreamId != 2 {
				var item models.Breaks
				item = brs
				stream_id, affiliate_id := findClick(brs.VCode)
				item.IsBreaked = 1
				item.AffiliateID = affiliate_id
				item.StreamId = stream_id
				item.Version = 1
				breaks_array = append(breaks_array, item)
			}
		} else if brks, ok := breakings[v.VCode]; ok {
			if brks.AffiliateID == 0 || brs.StreamId != 2 {
				var item models.Breaks
				item = brks
				stream_id, affiliate_id := findClick(brks.VCode)
				item.AffiliateID = affiliate_id
				item.StreamId = stream_id
				item.Version = 1
				breaks_array = append(breaks_array, item)
			}
		} else if cls, ok := clicks[v.VCode]; ok {
			var item models.Breaks
			stream := findFromMongo(cls.Campaign, cls.SourceID)
			item.VCode = cls.VCode
			item.CreateAt = cls.CreateAt
			item.Create_date = types.Date(cls.CreateAt)
			item.IsRefused = 0
			item.IsBreaked = 1
			item.AffiliateID = int32(stream[len(stream)-1].AffiliateID)
			item.StreamId = int32(stream[len(stream)-1 ].ID)
			item.Version = 1
			breaks_array = append(breaks_array, item)
		} else {
			fmt.Println("Клик с vcode=", v.VCode, " не найден")
		}
	}
	fmt.Println(breaks_array)
	if len(breaks_array) > 0 {
		fmt.Println("Запись пробива")
		query := `
					INSERT INTO tracker_db.breaks
						(vcode,
						process_interval,
						create_at,
						stream_id,
						is_breaked,
						affiliate_id, 
						screen_width, 
						screen_height,
						language,
						is_refused,
						version)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		clickhouse := database.SqlxConnect()

		tx, err := clickhouse.Begin()
		ErrorCheck(err)

		stmt, err := tx.Prepare(query)
		ErrorCheck(err)
		for _, item := range breaks_array {
			if _, err := stmt.Exec(
				item.VCode,
				item.ProcessInterval,
				item.CreateAt,
				item.StreamId,
				item.IsBreaked,
				item.AffiliateID,
				item.ScreenHeight,
				item.ScreenWidth,
				item.Language,
				item.IsRefused,
				1,
			); err != nil {
				fmt.Println(err.Error())
			}

		}
		if err := tx.Commit(); err != nil {
			fmt.Println(err.Error())
		}

		stmt.Close()
		clickhouse.Close()
	}
	fmt.Println("Всё записали")

}

func findClick(vcode string) (stream_id, affiliate_id int32) {
	if cls, ok := clicks[vcode]; ok {
		item := findFromMongo(cls.Campaign, cls.SourceID)
		return int32(item[len(item)-1].ID), int32(item[len(item)-1].AffiliateID)
	} else {
		return 0, 0
	}

}

func findFromMongo(company_id int32, source_id int32) []models.Stream {
	collection := database.MongoDB().C("stream")
	//var result []company

	result := []models.Stream{}
	// Limit(100).
	iter := collection.Find(bson.D{{"company_id", company_id}, {"source_id", 2}, {"type", 2}, {"is_archiving", false}}).Iter()
	err := iter.All(&result)

	if err = iter.Close(); err != nil {
		panic(err)
	}
	return result
}

func getBreaks(keys []string) {
	from := 0
	to := 3000
	max := len(keys)
	for {
		if to > max {
			to = max
		}
		if from == max {
			break
		}
		vcodeArray := keys[from:to]
		vcodeString := ""
		if len(vcodeArray) > 1 {
			vcodeString = "'" + strings.Join(vcodeArray, "','") + "'"
		} else {
			vcodeString = "'" + vcodeArray[0] + "'"
		}
		query := fmt.Sprintf(`
		SELECT * FROM tracker_db.breaks FINAL WHERE vcode in (%s)`, vcodeString)

		clickhouse := database.SqlxConnect()
		var collected_data []models.Breaks
		if err := clickhouse.Select(&collected_data, query); err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second)
		if len(collected_data) == 0 {
			break
		}

		for _, v := range collected_data {
			breaks[v.VCode] = v
		}
		from = to
		to = to + 3000
	}
	fmt.Println("Получили breaks")
}

func getBreakings(keys []string) {
	from := 0
	to := 3000
	max := len(keys)
	for {
		if to > max {
			to = max
		}
		if from == max {
			break
		}
		vcodeArray := keys[from:to]
		vcodeString := ""
		if len(vcodeArray) > 1 {
			vcodeString = "'" + strings.Join(vcodeArray, "','") + "'"
		} else {
			vcodeString = "'" + vcodeArray[0] + "'"
		}
		query := fmt.Sprintf(`
		SELECT * FROM tracker_db.breakings WHERE vcode in (%s)`, vcodeString)

		clickhouse := database.SqlxConnect()
		var collected_data []models.Breaks
		if err := clickhouse.Select(&collected_data, query); err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second)
		if len(collected_data) == 0 {
			break
		}

		for _, v := range collected_data {
			breakings[v.VCode] = v
		}
		from = to
		to = to + 3000
	}
	fmt.Println("Получили breakings")
}

func getClicks(keys []string) {
	from := 0
	to := 3000
	max := len(keys)
	for {
		if to > max {
			to = max
		}
		if from == max {
			break
		}
		vcodeArray := keys[from:to]
		vcodeString := ""
		if len(vcodeArray) > 1 {
			vcodeString = "'" + strings.Join(vcodeArray, "','") + "'"
		} else {
			vcodeString = "'" + vcodeArray[0] + "'"
		}
		query := fmt.Sprintf(`
		SELECT * FROM tracker_db.click_logs WHERE vcode in (%s) AND is_test = 0`, vcodeString)

		clickhouse := database.SqlxConnect()
		var collected_data []models.Click
		if err := clickhouse.Select(&collected_data, query); err != nil {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second)
		if len(collected_data) == 0 {
			break
		}

		for _, v := range collected_data {
			clicks[v.VCode] = v
		}
		from = to
		to = to + 3000
	}
	fmt.Println("Получили clicks")
}

func getLeads() []string {
	var leads_keys []string
	query := `
		SELECT * FROM tracker_db.leads final WHERE vcode != ''
		`

	clickhouse := database.SqlxConnect()
	var collected_data []models.PostBack
	if err := clickhouse.Select(&collected_data, query); err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second)
	for _, v := range collected_data {
		leads[v.VCode] = v
		leads_keys = append(leads_keys, v.VCode)
	}
	fmt.Println("Получили leads")
	return leads_keys
}

func AffiliateProgramList() []models.AffiliateProgram {
	collection := database.MongoDB().C("affiliate_program")
	//var result []company

	result := []models.AffiliateProgram{}
	// Limit(100).
	iter := collection.Find(nil).Iter()
	err := iter.All(&result)

	if err = iter.Close(); err != nil {
		panic(err)
	}
	return result
}

func rewriteToBreaks(affiliate []models.AffiliateProgram) {
	i := 0
	var ids_slice []int
	for _, v := range affiliate {
		ids_slice = append(ids_slice, v.AffiliateId)
	}
	for {
		query := fmt.Sprintf(`
		SELECT * FROM tracker_db.breakings
		ORDER BY create_at DESC limit %d,3000
		`, i)

		clickhouse := database.SqlxConnect()
		var collected_data []models.Breaks
		if err := clickhouse.Select(&collected_data, query); err != nil {
			fmt.Println(err)
			break
		}
		clickhouse.Close()
		time.Sleep(time.Second)
		i = i + 3000

		if len(collected_data) > 0 {
			fmt.Println("Запись пробива")
			query := `
					INSERT INTO tracker_db.breaks
						(vcode,
						process_interval,
						create_at,
						stream_id,
						is_breaked,
						affiliate_id, 
						screen_width, 
						screen_height,
						language,
						is_refused,
						version)
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

			clickhouse := database.SqlxConnect()

			tx, err := clickhouse.Begin()
			ErrorCheck(err)

			stmt, err := tx.Prepare(query)
			ErrorCheck(err)
			for _, item := range collected_data {
				if _, err := stmt.Exec(
					item.VCode,
					item.ProcessInterval,
					item.CreateAt,
					item.StreamId,
					item.IsBreaked,
					ids_slice[rand.Intn(len(ids_slice))],
					item.ScreenHeight,
					item.ScreenWidth,
					item.Language,
					item.IsRefused,
					1,
				); err != nil {
					fmt.Println(err.Error())
				}

			}
			if err := tx.Commit(); err != nil {
				fmt.Println(err.Error())
			}

			stmt.Close()
			clickhouse.Close()
		} else {
			fmt.Println("Данные закончились")
			break
		}

	}
}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
