package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"dataCatcher/db"
	"dataCatcher/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ClickhouseInit()
	rewriteData()
}

func rewriteData(){
	query := `
		SELECT
				vcode,
				click_logs.create_at as create_at,
				create_date,
				url,
				method, 
				params, 
				status_confirmed,
				status_hold,
				status_declined,
				status_other,
				order_id,
				amount,
				result_message
		FROM tracker_db.leads
		ALL LEFT JOIN tracker_db.click_logs USING vcode
		WHERE toDate(click_logs.create_at) <> toDate(leads.create_date) AND click_logs.create_at != 0`

	clickhouse := database.SqlxConnect()
	var collected_data []models.PostBack
	if err := clickhouse.Select(&collected_data, query); err != nil {
		fmt.Println(err)
	}
	clickhouse.Close()



	update_query := fmt.Sprintf(`
				INSERT INTO tracker_db.leads
					(vcode,
					create_at,
					create_date,
					url,
					method, 
					params, 
					status_confirmed,
					status_hold,
					status_declined,
					status_other,
					order_id,
					amount,
					profit,
					result_message,
					version)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	clickhouse = database.SqlxConnect()
	tx, err := clickhouse.Begin()
	ErrorCheck(err)

	stmt, err := tx.Prepare(update_query)
	ErrorCheck(err)

	for _, item := range collected_data {
		if _, err := stmt.Exec(
			item.VCode,
			item.CreateAt,
			item.CreateDate,
			item.Url,
			item.Method,
			item.Params,
			item.StatusConfirmed,
			item.StatusHold,
			item.StatusDeclined,
			item.StatusOther,
			item.OrderID,
			item.Amount,
			item.Profit,
			item.ResultMessage,
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

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}