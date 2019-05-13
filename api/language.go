package api

import (
	"net/http"

	"generated/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configLanguagesRouter(router *httprouter.Router) {
	router.GET("/languages", GetAllLanguages)
	router.POST("/languages", AddLanguage)
	router.GET("/languages/:id", GetLanguage)
	router.PUT("/languages/:id", UpdateLanguage)
	router.DELETE("/languages/:id", DeleteLanguage)
}

func GetAllLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	languages := []*model.Language{}

	if order != "" {
		err = DB.Model(&model.Language{}).Order(order).Offset(offset).Limit(pagesize).Find(&languages).Error
	} else {
		err = DB.Model(&model.Language{}).Offset(offset).Limit(pagesize).Find(&languages).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetLanguage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	language := &model.Language{}
	if DB.First(language, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, language)
}

func AddLanguage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	language := &model.Language{}

	if err := readJSON(r, language); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(language).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, language)
}

func UpdateLanguage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	language := &model.Language{}
	if DB.First(language, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Language{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(language, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(language).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, language)
}

func DeleteLanguage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	language := &model.Language{}

	if DB.First(language, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(language).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
