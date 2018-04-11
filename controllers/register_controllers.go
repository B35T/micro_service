package controllers

import (
	"net/http"
	"micro_service/models"
	"fmt"
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

