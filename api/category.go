package api

import (
	"net/http"

	"generated/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configCategoriesRouter(router *httprouter.Router) {
	router.GET("/categories", GetAllCategories)
	router.POST("/categories", AddCategory)
	router.GET("/categories/:id", GetCategory)
	router.PUT("/categories/:id", UpdateCategory)
	router.DELETE("/categories/:id", DeleteCategory)
}

func GetAllCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	categories := []*model.Category{}

	if order != "" {
		err = DB.Model(&model.Category{}).Order(order).Offset(offset).Limit(pagesize).Find(&categories).Error
	} else {
		err = DB.Model(&model.Category{}).Offset(offset).Limit(pagesize).Find(&categories).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	category := &model.Category{}
	if DB.First(category, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, category)
}

func AddCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := &model.Category{}

	if err := readJSON(r, category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	category := &model.Category{}
	if DB.First(category, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Category{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(category, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	category := &model.Category{}

	if DB.First(category, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
