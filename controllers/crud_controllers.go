package controllers

import (
	"micro_service/db"
	"fmt"
	"micro_service/models"
)

func InsertUser(u *models.PersonStruct) (bool, error ) {
	var db = db.NewDB().DB
	defer db.Close()

	sql := fmt.Sprintf("INSERT INTO Persons (PersonID, LastName, FirstName, Address, City) " +
		"VALUES ('%s', '%s', '%s', '%s', '%s');" ,u.PersonID, u.LastName, u.FirstName, u.Address, u.City)

	_, err := db.Exec(sql)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
