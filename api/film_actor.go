package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configFilmActorsRouter(router *httprouter.Router) {
	router.GET("/filmactors", GetAllFilmActors)
	router.POST("/filmactors", AddFilmActor)
	router.GET("/filmactors/:id", GetFilmActor)
	router.PUT("/filmactors/:id", UpdateFilmActor)
	router.DELETE("/filmactors/:id", DeleteFilmActor)
}

func GetAllFilmActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	filmactors := []*model.FilmActor{}

	if order != "" {
		err = DB.Model(&model.FilmActor{}).Order(order).Offset(offset).Limit(pagesize).Find(&filmactors).Error
	} else {
		err = DB.Model(&model.FilmActor{}).Offset(offset).Limit(pagesize).Find(&filmactors).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetFilmActor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	filmactor := &model.FilmActor{}
	if DB.First(filmactor, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, filmactor)
}

func AddFilmActor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filmactor := &model.FilmActor{}

	if err := readJSON(r, filmactor); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(filmactor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, filmactor)
}

func UpdateFilmActor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	filmactor := &model.FilmActor{}
	if DB.First(filmactor, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.FilmActor{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(filmactor, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(filmactor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, filmactor)
}

func DeleteFilmActor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	filmactor := &model.FilmActor{}

	if DB.First(filmactor, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(filmactor).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
