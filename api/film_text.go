package api

import (
	"net/http"

	"generated/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configFilmTextsRouter(router *httprouter.Router) {
	router.GET("/filmtexts", GetAllFilmTexts)
	router.POST("/filmtexts", AddFilmText)
	router.GET("/filmtexts/:id", GetFilmText)
	router.PUT("/filmtexts/:id", UpdateFilmText)
	router.DELETE("/filmtexts/:id", DeleteFilmText)
}

func GetAllFilmTexts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	filmtexts := []*model.FilmText{}

	if order != "" {
		err = DB.Model(&model.FilmText{}).Order(order).Offset(offset).Limit(pagesize).Find(&filmtexts).Error
	} else {
		err = DB.Model(&model.FilmText{}).Offset(offset).Limit(pagesize).Find(&filmtexts).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetFilmText(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	filmtext := &model.FilmText{}
	if DB.First(filmtext, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, filmtext)
}

func AddFilmText(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filmtext := &model.FilmText{}

	if err := readJSON(r, filmtext); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(filmtext).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, filmtext)
}

func UpdateFilmText(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	filmtext := &model.FilmText{}
	if DB.First(filmtext, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.FilmText{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(filmtext, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(filmtext).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, filmtext)
}

func DeleteFilmText(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	filmtext := &model.FilmText{}

	if DB.First(filmtext, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(filmtext).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
