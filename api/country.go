package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configCountriesRouter(router *httprouter.Router) {
	router.GET("/countries", GetAllCountries)
	router.POST("/countries", AddCountry)
	router.GET("/countries/:id", GetCountry)
	router.PUT("/countries/:id", UpdateCountry)
	router.DELETE("/countries/:id", DeleteCountry)
}

func GetAllCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	countries := []*model.Country{}

	if order != "" {
		err = DB.Model(&model.Country{}).Order(order).Offset(offset).Limit(pagesize).Find(&countries).Error
	} else {
		err = DB.Model(&model.Country{}).Offset(offset).Limit(pagesize).Find(&countries).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCountry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	country := &model.Country{}
	if DB.First(country, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, country)
}

func AddCountry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	country := &model.Country{}

	if err := readJSON(r, country); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(country).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, country)
}

func UpdateCountry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	country := &model.Country{}
	if DB.First(country, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Country{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(country, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(country).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, country)
}

func DeleteCountry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	country := &model.Country{}

	if DB.First(country, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(country).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
