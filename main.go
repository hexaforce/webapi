package main

import (
	"log"
	"net/http"

	"github.com/hexaforce/webapi/api"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	DB, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sakila?charset=utf8&parseTime=true")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	defer DB.Close()

	router := api.ConfigRouter(DB)
	log.Fatal(http.ListenAndServe(":8080", router))

}
