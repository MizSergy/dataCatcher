package controllers

import (
	"fmt"
	"redisCatcher/db"
	"redisCatcher/models"
	"strconv"
	"time"
)

func Ban(){
	listArray := []int{
		419,
		3973,
		3974,
		3975,
		3976,
		3977,
		3978,
		3979,
		3980,
		3981,
		3982,
		3983,
		3984,
		3985,
		3986,
		3987,
		4163,
		4253,
		4254,
		4941,
		4942,
		4943,
		4944,
		4945,
		4991,
		5024,
		5025,
		5026,
		5140,
		5141,
		5142,
		5143,
		5244,
		5245,
		5246,
		5247,
		5248,
		5249,
		5651,
		5652,
		5654,
		5824,
		5825,
		5827,
		5828,
		5829,
		8504,
		10048,
		10050,
		10052,
		10054,
		10056,
		10058,
		10060,
		10062,
		10064,
		10066,
		10068,
		10070,
		10072,
		10074,
		10076,
		10078,
		10080,
		10082,
		10084,
		10086,
		10088,
		10090,
		10092,
		10094,
		10096,
		10098,
		10100,
		10102,
		10104,
		10106,
		10108,
		10110,
		10112,
		10114,
		10116,
		10118,
		10120,
		10122,
		10124,
		10126,
		10659,
		10660,
		10661,
		10662,
		10663,
		10664,
		10665,
		10666,
		10667,
		10668,
		10669,
		10670,
		10671,
		10672,
		10673,
		10674,
		10675,
		10676,
		10677,
		10678}

	var model_array []models.BlackLists

	for i := 0; i < len(listArray); i++ {
		model := models.BlackLists{
			Site : strconv.Itoa(listArray[i]),
			SourceID : 2,
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
