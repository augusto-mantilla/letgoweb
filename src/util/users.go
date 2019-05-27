package util

import (
	"fmt"
	"log"
	"strconv"
)

type Address struct {
	ZIP_code   string
	Ph_address string
}

type User struct {
	name         string
	nif          string
	phone_number string
	email        string
	foto_url     string
	address_id   int
}

func InsertAddress(zipcode string, phaddress string) {
	query := "INSERT INTO address (zip_code, address) VALUES ('" + zipcode + "', '" + phaddress + "')"
	row, err := db.Query(query)
	if err != nil {
		log.Panic(err.Error())
	}

	fmt.Println(row)
}

func InsertUser(name, nif, phone_number, email, foto_url string, address_id int) {
	query := "INSERT INTO user (name, nif, phone_number, email, foto_url, address_id)"
	query += "VALUES (" + name + ", " + nif + ", " + phone_number + ", " + email +
		", " + foto_url + ", " + strconv.Itoa(address_id) + ")"
	db.Query(query)
}
