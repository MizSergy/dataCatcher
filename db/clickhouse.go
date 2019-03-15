package database

import (
"fmt"
"github.com/jmoiron/sqlx"
_ "github.com/kshvakov/clickhouse"
"os"
)

var ListClickhouseConnect []ClickhouseClient

//--------------------------------------------------Новое подключение кликхауса-----------------------------------------

type ClickhouseClient struct {
	sqlxclient *sqlx.DB
	Host       string
	Port       string
	Username   string
	Password   string
	DB         string
	Debug      string
}

func ClickhouseInit() {
	index := 1
	for {
		host := os.Getenv(fmt.Sprintf("CLICKHOUSE_HOST_%d", index))
		if len(host) == 0 {
			break
		}
		port := os.Getenv(fmt.Sprintf("CLICKHOUSE_PORT_%d", index))
		username := os.Getenv(fmt.Sprintf("CLICKHOUSE_USERNAME_%d", index))
		pass := os.Getenv(fmt.Sprintf("CLICKHOUSE_PASS_%d", index))
		db := os.Getenv(fmt.Sprintf("CLICKHOUSE_DB_%d", index))

		ListClickhouseConnect = append(ListClickhouseConnect, ClickhouseClient{
			Host:     host,
			Port:     port,
			Username: username,
			Password: pass,
			DB:       db,
			Debug:    os.Getenv("CLICKHOUSE_DEBUG"),
		})
		fmt.Printf("Connect clickhouse %s:%s\n", host, port)
		index ++
	}
	fmt.Printf("Clickhouse count connect %d\n", len(ListClickhouseConnect))

}

func (cl *ClickhouseClient) Disconnect() {
	cl.sqlxclient.Close()
}

func SqlxConnect() *sqlx.DB {
	var connection *sqlx.DB
	for _,v := range ListClickhouseConnect{
		if v.sqlxclient == nil{
			params := fmt.Sprintf(`%s:%s?username=%s&password=%s&database=%s&debug=%s`, v.Host, v.Port, v.Username, v.Password, v.DB, v.Debug)
			var err error
			v.sqlxclient, err = sqlx.Open("clickhouse", params)
			if err != nil {
				fmt.Printf("%s:%s - no connect\n", v.Host, v.Port)
				return nil
			}
			connection = v.sqlxclient
		}
		connection = v.sqlxclient
	}
	return connection
}
