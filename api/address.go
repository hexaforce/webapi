package api

import (
	"net/http"

	"generated/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configAddressesRouter(router *httprouter.Router) {
	router.GET("/addresses", GetAllAddresses)
	router.POST("/addresses", AddAddress)
	router.GET("/addresses/:id", GetAddress)
	router.PUT("/addresses/:id", UpdateAddress)
	router.DELETE("/addresses/:id", DeleteAddress)
}

func GetAllAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	addresses := []*model.Address{}

	if order != "" {
		err = DB.Model(&model.Address{}).Order(order).Offset(offset).Limit(pagesize).Find(&addresses).Error
	} else {
		err = DB.Model(&model.Address{}).Offset(offset).Limit(pagesize).Find(&addresses).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	address := &model.Address{}
	if DB.First(address, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, address)
}

func AddAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	address := &model.Address{}

	if err := readJSON(r, address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(address).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, address)
}

func UpdateAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	address := &model.Address{}
	if DB.First(address, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Address{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(address, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(address).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, address)
}

func DeleteAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	address := &model.Address{}

	if DB.First(address, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(address).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
