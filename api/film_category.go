package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configFilmCategoriesRouter(router *httprouter.Router) {
	router.GET("/filmcategories", GetAllFilmCategories)
	router.POST("/filmcategories", AddFilmCategory)
	router.GET("/filmcategories/:id", GetFilmCategory)
	router.PUT("/filmcategories/:id", UpdateFilmCategory)
	router.DELETE("/filmcategories/:id", DeleteFilmCategory)
}

func GetAllFilmCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	filmcategories := []*model.FilmCategory{}

	if order != "" {
		err = DB.Model(&model.FilmCategory{}).Order(order).Offset(offset).Limit(pagesize).Find(&filmcategories).Error
	} else {
		err = DB.Model(&model.FilmCategory{}).Offset(offset).Limit(pagesize).Find(&filmcategories).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetFilmCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	filmcategory := &model.FilmCategory{}
	if DB.First(filmcategory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, filmcategory)
}

func AddFilmCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	filmcategory := &model.FilmCategory{}

	if err := readJSON(r, filmcategory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(filmcategory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, filmcategory)
}

func UpdateFilmCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	filmcategory := &model.FilmCategory{}
	if DB.First(filmcategory, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.FilmCategory{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(filmcategory, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(filmcategory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, filmcategory)
}

func DeleteFilmCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	filmcategory := &model.FilmCategory{}

	if DB.First(filmcategory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(filmcategory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
