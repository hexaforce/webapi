package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configCitiesRouter(router *httprouter.Router) {
	router.GET("/cities", GetAllCities)
	router.POST("/cities", AddCity)
	router.GET("/cities/:id", GetCity)
	router.PUT("/cities/:id", UpdateCity)
	router.DELETE("/cities/:id", DeleteCity)
}

func GetAllCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	page, err := readInt(r, "page", 1)
	if err != nil || page < 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	offset := (page - 1) * pagesize

	order := r.FormValue("order")

	cities := []*model.City{}

	if order != "" {
		err = DB.Model(&model.City{}).Order(order).Offset(offset).Limit(pagesize).Find(&cities).Error
	} else {
		err = DB.Model(&model.City{}).Offset(offset).Limit(pagesize).Find(&cities).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	city := &model.City{}
	if DB.First(city, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, city)
}

func AddCity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	city := &model.City{}

	if err := readJSON(r, city); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(city).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, city)
}

func UpdateCity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	city := &model.City{}
	if DB.First(city, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.City{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(city, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(city).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, city)
}

func DeleteCity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	city := &model.City{}

	if DB.First(city, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(city).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
