package controllers

import (
	"fmt"
	"redisCatcher/db"
	"redisCatcher/models"
)

func FillClicks(){
	select_query := fmt.Sprintf(`SELECT 
		 	create_at,
			vcode,
			is_unique,
			campaign,
			source_id, 
			br.affiliate_id as affiliate_id, 
			click_price, 
			is_mobil,
			device,
			browser,
			os,
			country,
			region,
			city,
			ip,
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
			preland_id,
			session_id,
			is_test,
			country_code,
			osv,
			browserv
FROM tracker_db.click_logs
ALL LEFT JOIN tracker_db.breaks as br FINAL USING vcode
PREWHERE toDate(create_at) BETWEEN '2019-03-05' and '2019-03-15'`)

	clickhouse := database.SqlxConnect()
	var collected_data []models.Click
	if err := clickhouse.Select(&collected_data, select_query); err != nil {
		fmt.Println(err)
	}
	clickhouse.Close()
	fmt.Println("Взяли")

	if (len(collected_data) > 0) {
		fmt.Println("Запись кликов")

		query := `INSERT INTO tracker_db.click_logs
			(create_at,
			vcode,
			is_unique,
			campaign,
			source_id, 
			affiliate_id, 
			click_price, 
			is_mobil,
			device,
			browser,
			os,
			country,
			region,
			city,
			ip,
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
			preland_id,
			session_id,
			is_test,
			country_code,
			osv,
			browserv)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

		clickhouse_conn := database.SqlxConnect()

		tx, err := clickhouse_conn.Begin()
		ErrorCheck(err)

		stmt, err := tx.Prepare(query)
		ErrorCheck(err)

		for _, item := range collected_data {
			if _, err := stmt.Exec(
				item.CreateAt,
				item.VCode,
				item.IsUnique,
				item.Campaign,
				item.SourceID,
				item.AffiliateID,
				item.ClickPrice,
				item.IsMobil,
				item.Device,
				item.Browser,
				item.Os,
				item.Country,
				item.Region,
				item.City,
				item.Ip,
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
				item.PrelandID,
				item.Session,
				item.IsTest,
				item.CountryCode,
				item.OsV,
				item.BrowserV,
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
}