package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"micro_service/config"
)


var DB = NewDB()

type PgStore struct {
	DB *sql.DB
}

func NewDB() *PgStore {
	var conf = config.Open_file

	//conf := Open_file
	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DB_Name, conf.DB_Pw, conf.DB)

	//"postgres://user:password@ip:port/dbname?sslmode=disable"
	// example "postgres://admin:root@192.168.1.5:5432/users_store?sslmode=disable"
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DB_user,conf.DB_Pw,conf.DB_Ip,conf.DB_port,conf.DB)
	db, err := sql.Open("postgres",url)
	if err != nil {
		panic(err)
	}

	return &PgStore{db}
}

