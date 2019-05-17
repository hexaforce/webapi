package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configRentalsRouter(router *httprouter.Router) {
	router.GET("/rentals", GetAllRentals)
	router.POST("/rentals", AddRental)
	router.GET("/rentals/:id", GetRental)
	router.PUT("/rentals/:id", UpdateRental)
	router.DELETE("/rentals/:id", DeleteRental)
}

func GetAllRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	rentals := []*model.Rental{}

	if order != "" {
		err = DB.Model(&model.Rental{}).Order(order).Offset(offset).Limit(pagesize).Find(&rentals).Error
	} else {
		err = DB.Model(&model.Rental{}).Offset(offset).Limit(pagesize).Find(&rentals).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetRental(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	rental := &model.Rental{}
	if DB.First(rental, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, rental)
}

func AddRental(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rental := &model.Rental{}

	if err := readJSON(r, rental); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(rental).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, rental)
}

func UpdateRental(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	rental := &model.Rental{}
	if DB.First(rental, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Rental{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(rental, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(rental).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, rental)
}

func DeleteRental(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	rental := &model.Rental{}

	if DB.First(rental, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(rental).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
