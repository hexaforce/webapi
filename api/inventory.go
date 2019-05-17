package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configInventoriesRouter(router *httprouter.Router) {
	router.GET("/inventories", GetAllInventories)
	router.POST("/inventories", AddInventory)
	router.GET("/inventories/:id", GetInventory)
	router.PUT("/inventories/:id", UpdateInventory)
	router.DELETE("/inventories/:id", DeleteInventory)
}

func GetAllInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	inventories := []*model.Inventory{}

	if order != "" {
		err = DB.Model(&model.Inventory{}).Order(order).Offset(offset).Limit(pagesize).Find(&inventories).Error
	} else {
		err = DB.Model(&model.Inventory{}).Offset(offset).Limit(pagesize).Find(&inventories).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	inventory := &model.Inventory{}
	if DB.First(inventory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, inventory)
}

func AddInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inventory := &model.Inventory{}

	if err := readJSON(r, inventory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(inventory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, inventory)
}

func UpdateInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	inventory := &model.Inventory{}
	if DB.First(inventory, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Inventory{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(inventory, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(inventory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, inventory)
}

func DeleteInventory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	inventory := &model.Inventory{}

	if DB.First(inventory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(inventory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
