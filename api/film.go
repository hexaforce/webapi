package api

import (
	"net/http"

	"generated/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configFilmsRouter(router *httprouter.Router) {
	router.GET("/films", GetAllFilms)
	router.POST("/films", AddFilm)
	router.GET("/films/:id", GetFilm)
	router.PUT("/films/:id", UpdateFilm)
	router.DELETE("/films/:id", DeleteFilm)
}

func GetAllFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	films := []*model.Film{}

	if order != "" {
		err = DB.Model(&model.Film{}).Order(order).Offset(offset).Limit(pagesize).Find(&films).Error
	} else {
		err = DB.Model(&model.Film{}).Offset(offset).Limit(pagesize).Find(&films).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetFilm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	film := &model.Film{}
	if DB.First(film, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, film)
}

func AddFilm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	film := &model.Film{}

	if err := readJSON(r, film); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(film).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, film)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	film := &model.Film{}
	if DB.First(film, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Film{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(film, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(film).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, film)
}

func DeleteFilm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	film := &model.Film{}

	if DB.First(film, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(film).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
