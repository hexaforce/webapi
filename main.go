package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

func main() {

	DB, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sakila?charset=utf8&parseTime=true")
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	defer db.Close()

    router := api.ConfigRouter(DB)
	log.Fatal(http.ListenAndServe(":8080", router))
	
}
