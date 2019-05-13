package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

var DB *gorm.DB

func ConfigRouter(DB *gorm.DB) http.Handler {
	router := httprouter.New()
	configActorsRouter(router)
	configAddressesRouter(router)
	configCategoriesRouter(router)
	configCitiesRouter(router)
	configCountriesRouter(router)
	configCustomersRouter(router)
	configFilmsRouter(router)
	configFilmActorsRouter(router)
	configFilmCategoriesRouter(router)
	configFilmTextsRouter(router)
	configInventoriesRouter(router)
	configLanguagesRouter(router)
	configPaymentsRouter(router)
	configRentalsRouter(router)
	configStaffsRouter(router)
	configStoresRouter(router)

	return router
}

func readInt(r *http.Request, param string, v int64) (int64, error) {
	p := r.FormValue(param)
	if p == "" {
		return v, nil
	}
	return strconv.ParseInt(p, 10, 64)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
