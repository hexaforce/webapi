package api

import (
	"net/http"

	"github.com/hexaforce/webapi/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configStaffsRouter(router *httprouter.Router) {
	router.GET("/staffs", GetAllStaffs)
	router.POST("/staffs", AddStaff)
	router.GET("/staffs/:id", GetStaff)
	router.PUT("/staffs/:id", UpdateStaff)
	router.DELETE("/staffs/:id", DeleteStaff)
}

func GetAllStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	staffs := []*model.Staff{}

	if order != "" {
		err = DB.Model(&model.Staff{}).Order(order).Offset(offset).Limit(pagesize).Find(&staffs).Error
	} else {
		err = DB.Model(&model.Staff{}).Offset(offset).Limit(pagesize).Find(&staffs).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetStaff(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	staff := &model.Staff{}
	if DB.First(staff, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, staff)
}

func AddStaff(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staff := &model.Staff{}

	if err := readJSON(r, staff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(staff).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, staff)
}

func UpdateStaff(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	staff := &model.Staff{}
	if DB.First(staff, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Staff{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(staff, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(staff).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, staff)
}

func DeleteStaff(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	staff := &model.Staff{}

	if DB.First(staff, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(staff).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
