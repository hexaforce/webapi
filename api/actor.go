package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configActorsRouter(router *httprouter.Router) {
	router.GET("/actors", GetAllActors)
	router.POST("/actors", AddActor)
	router.GET("/actors/:id", GetActor)
	router.PUT("/actors/:id", UpdateActor)
	router.DELETE("/actors/:id", DeleteActor)
}

func GetAllActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	actors := []*model.Actor{}

	if order != "" {
		err = DB.Model(&model.Actor{}).Order(order).Offset(offset).Limit(pagesize).Find(&actors).Error
	} else {
		err = DB.Model(&model.Actor{}).Offset(offset).Limit(pagesize).Find(&actors).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetActor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	actor := &model.Actor{}
	if DB.First(actor, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, actor)
}

func AddActor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	actor := &model.Actor{}

	if err := readJSON(r, actor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(actor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, actor)
}

func UpdateActor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	actor := &model.Actor{}
	if DB.First(actor, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Actor{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(actor, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(actor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, actor)
}

func DeleteActor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	actor := &model.Actor{}

	if DB.First(actor, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(actor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
