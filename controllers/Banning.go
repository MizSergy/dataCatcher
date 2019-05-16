package controllers

import (
	"fmt"
	"redisCatcher/db"
	"redisCatcher/models"
	"time"
)

func Ban(){
	listArray := []string{
		"18302",
		"17933",
		"17845",
		"17666",
		"17034",
		"16611",
		"16610",
		"16608",
		"16419",
		"15772",
		"15644",
		"14951",
		"14279",
		"13870",
		"13869",
		"11884",
		"11711",
		"10507",
		"10461",
		"7152",
		"5731",
		"3744",
		"3120",
		"2275",
		"2273",
		"2194",
		"1833",
		"669",
		"335",
		"257",
		"61",
		"1",
		"29",
		"7",
		"4",
		"30"}

	var model_array []models.BlackLists

	for i := 0; i < len(listArray); i++ {
		model := models.BlackLists{
			Site : listArray[i],
			SourceID : 10,
			CompanyID : 1,
			Ban : 1,
		}
		model.Version = 1
		model.CreateAt = time.Now()
		model_array = append(model_array, model)
	}
	query := `
					INSERT INTO tracker_db.blacklists
						(site,
						source_id,
						campaign,
						version,
						ban)
					VALUES (?, ?, ?, ?)`

	clickhouse := database.SqlxConnect()

	tx, err := clickhouse.Begin()

	ErrorCheck(err)

	stmt, err := tx.Prepare(query)
	ErrorCheck(err)
	for _, item := range model_array {
		if _, err := stmt.Exec(
			item.Site,
			item.SourceID,
			item.CompanyID,
			item.Version,
			item.Ban,
		); err != nil {
			fmt.Println(err.Error())
		}

	}
	if err := tx.Commit(); err != nil {
		fmt.Println(err.Error())
	}
	stmt.Close()
	clickhouse.Close()
	fmt.Println("Всё")

}


func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
