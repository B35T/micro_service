package controllers

import (
	"net/http"
	"micro_service/models"
	"fmt"
	"encoding/json"
)

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		data := &models.PersonStruct{
			PersonID:r.FormValue("person_id"),
			FirstName:r.FormValue("firstname"),
			LastName:r.FormValue("lastname"),
			Address:r.FormValue("address"),
			City:r.FormValue("city"),
		}

		_, err := InsertUser(data)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("insert data successful")
		}
		break
	}
}

func JsonAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p := &models.PersonStruct{
			PersonID:"000000000",
			FirstName:"chaloemphong",
			LastName:"chaimool",
			Address:"1/1",
			City:"SSK",
		}

		j, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w,string(j))
		break
	}
}

func TestConnect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"test connect")
}


