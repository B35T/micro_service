package config

import (
	"io/ioutil"
	"encoding/json"
)

type config struct {
	DB string `json:"db"`
	DB_user string `json:"db_user"`
	DB_Pw string `json:"db_pw"`
	DB_Ip string `json:"db_ip"`
	DB_port string `json:"db_port"`
	DB_Tb_Name string `json:"db_tb_name"`
}

var Open_file = func() *config {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}

	var c config
	err = json.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}

	return &c
}()
