package controllers

import (
	"fmt"
	"redisCatcher/db"
	"redisCatcher/models"
)

func CollectClicksBreaks() {
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
	var trafficData []models.FullTraffic
	for _,v := range collected_data{
		trafficData = append(trafficData, v.Merge(models.FullTraffic{}))
	}
}
