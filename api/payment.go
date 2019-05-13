package api

import (
	"net/http"

	"generated/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configPaymentsRouter(router *httprouter.Router) {
	router.GET("/payments", GetAllPayments)
	router.POST("/payments", AddPayment)
	router.GET("/payments/:id", GetPayment)
	router.PUT("/payments/:id", UpdatePayment)
	router.DELETE("/payments/:id", DeletePayment)
}

func GetAllPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	payments := []*model.Payment{}

	if order != "" {
		err = DB.Model(&model.Payment{}).Order(order).Offset(offset).Limit(pagesize).Find(&payments).Error
	} else {
		err = DB.Model(&model.Payment{}).Offset(offset).Limit(pagesize).Find(&payments).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetPayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	payment := &model.Payment{}
	if DB.First(payment, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, payment)
}

func AddPayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payment := &model.Payment{}

	if err := readJSON(r, payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(payment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, payment)
}

func UpdatePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	payment := &model.Payment{}
	if DB.First(payment, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Payment{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(payment, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(payment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, payment)
}

func DeletePayment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	payment := &model.Payment{}

	if DB.First(payment, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(payment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
