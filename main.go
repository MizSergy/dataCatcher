package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"redisCatcher/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.RedisInit()
	var check_time_data []string
	var break_data []string

	for _, client := range database.ListRedisConnect {
		if client.Connect() == nil {
			continue
		}
		check_time_data, err = client.Connect().LRange("checkTime",0,100000000000).Result()
		if err!= nil{
			fmt.Println(err.Error())
		}
		break_data, err = client.Connect().LRange("checkTime",0,100000000000).Result()
		if err!= nil{
			fmt.Println(err.Error())
		}
	}

	if len(break_data) > 0{
		fmt.Println("Checktime collected")
		//--------------------------------Запись данных в файл--------------------------
		f, err := os.Create("dump.txt")
		if err != nil {
			fmt.Println("Unable to create file:", err)
		}
		defer f.Close()

		data, err := json.Marshal(break_data)
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Println("Записываем данные в файл")
		f.Write(data)
	}

	if len(check_time_data) > 0{
		fmt.Println("Checktime collected")
		//--------------------------------Запись данных в файл--------------------------
		f, err := os.Create("dump.txt")
		if err != nil {
			fmt.Println("Unable to create file:", err)
		}
		defer f.Close()

		data, err := json.Marshal(check_time_data)
		if err != nil{
			fmt.Println(err.Error())
		}
		fmt.Println("Записываем данные в файл")
		f.Write(data)
	}

	//checktime_array := []models.Breaking{}
	//total_array := []models.Breaking{}
	//breaks_array := []models.Breaking{}
	//for _, client := range database.ListRedisConnect {
	//	if client.Connect() == nil {
	//		continue
	//	}
	//	//--------------------------------Собираем чектайм--------------------------
	//	for {
	//		check_time_data := client.GetQueueCollections("checkTime")
	//		if check_time_data == "" {
	//			break
	//		}
	//		breaks := models.Breaking{}
	//		err := json.Unmarshal([]byte(check_time_data), &breaks)
	//		if err != nil {
	//			log.Println(err.Error())
	//		}
	//		total_array = append(total_array, breaks)
	//		checktime_array = append(checktime_array, breaks)
	//		client.Connect().RPush("checkTime", breaks)
	//	}
	//
	//	fmt.Println("Checktime собран", checktime_array)
	//	//-----------------------------------Записали чектаймы----------------------
	//	if len(checktime_array) > 0{
	//		f, err := os.Create("checktime_dump.txt")
	//		if err != nil {
	//			fmt.Println("Unable to create file:", err)
	//		}
	//		var d []interface{}
	//
	//		for _, v := range checktime_array {
	//			d = append(d, v)
	//		}
	//		data, err := json.Marshal(d)
	//		if err != nil{
	//			fmt.Println(err.Error())
	//		}
	//		fmt.Println("Записываем checktime в файл")
	//		f.Write(data)
	//		f.Close()
	//	}
	//
	//	//--------------------------------Собираем пробивы--------------------------
	//	for{
	//		queue_breaking := client.GetQueueCollections("queue_breaking")
	//		if queue_breaking == "" {
	//			break
	//		}
	//		breaks := models.Breaking{}
	//		err := json.Unmarshal([]byte(queue_breaking), &breaks)
	//		if err != nil {
	//			log.Println(err.Error())
	//		}
	//		total_array = append(total_array, breaks)
	//		breaks_array = append(breaks_array, breaks)
	//		client.Connect().RPush("queue_breaking", breaks)
	//	}
	//}
	//if len(breaks_array) > 0{
	//	fmt.Println("Data collected")
	//	//--------------------------------Запись данных в файл--------------------------
	//	f, err := os.Create("dump.txt")
	//	if err != nil {
	//		fmt.Println("Unable to create file:", err)
	//	}
	//	var d []interface{}
	//	defer f.Close()
	//
	//	for _, v := range breaks_array {
	//		d = append(d, v)
	//	}
	//	data, err := json.Marshal(d)
	//	if err != nil{
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println("Записываем данные в файл")
	//	f.Write(data)
	//}
	//if len(breaks_array) > 0{
	//	fmt.Println("Breaks collected")
	//	//--------------------------------Запись данных в файл--------------------------
	//	f, err := os.Create("breacks_dump.txt")
	//	if err != nil {
	//		fmt.Println("Unable to create file:", err)
	//	}
	//	var d []interface{}
	//	defer f.Close()
	//
	//	for _, v := range breaks_array {
	//		d = append(d, v)
	//	}
	//	data, err := json.Marshal(d)
	//	if err != nil{
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println("Записываем данные в файл")
	//	f.Write(data)
	//}
	//
	//if len(total_array) > 0{
	//	fmt.Println("Total collected")
	//	//--------------------------------Запись данных в файл--------------------------
	//	f, err := os.Create("dump.txt")
	//	if err != nil {
	//		fmt.Println("Unable to create file:", err)
	//	}
	//	var d []interface{}
	//	defer f.Close()
	//
	//	for _, v := range breaks_array {
	//		d = append(d, v)
	//	}
	//	data, err := json.Marshal(d)
	//	if err != nil{
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println("Записываем данные в файл")
	//	f.Write(data)
	//}
}
