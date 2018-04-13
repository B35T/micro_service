package config

import (
	"io/ioutil"
	"encoding/json"
	"os/exec"
)

type config struct {
	DB string `json:"db"`
	DB_user string `json:"db_user"`
	DB_Pw string `json:"db_pw"`
	DB_Ip string `json:"db_ip"`
	DB_port string `json:"db_port"`
	DB_Tb_Name string `json:"db_tb_name"`
}

var urlConfig = func() string {
	out,_ := exec.Command("uname").Output()
	var s = string(out)
	darwin := string([]byte{68, 97, 114, 119, 105, 110, 10,})
	if s == darwin {
		return "src/micro_service/config/config.json"
	}
	return  "config/config.json"
}()

var Open_file = func() *config {
	data, err := ioutil.ReadFile(urlConfig)
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
