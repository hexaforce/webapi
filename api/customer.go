package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configCustomersRouter(router *httprouter.Router) {
	router.GET("/customers", GetAllCustomers)
	router.POST("/customers", AddCustomer)
	router.GET("/customers/:id", GetCustomer)
	router.PUT("/customers/:id", UpdateCustomer)
	router.DELETE("/customers/:id", DeleteCustomer)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	customers := []*model.Customer{}

	if order != "" {
		err = DB.Model(&model.Customer{}).Order(order).Offset(offset).Limit(pagesize).Find(&customers).Error
	} else {
		err = DB.Model(&model.Customer{}).Offset(offset).Limit(pagesize).Find(&customers).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCustomer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	customer := &model.Customer{}
	if DB.First(customer, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, customer)
}

func AddCustomer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customer := &model.Customer{}

	if err := readJSON(r, customer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(customer).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	customer := &model.Customer{}
	if DB.First(customer, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Customer{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(customer, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(customer).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	customer := &model.Customer{}

	if DB.First(customer, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(customer).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
