package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/binarydud/covidapi/db"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/hlog"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
func StateHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	name := chi.URLParam(r, "state")
	dbclient := db.New()
	state, err := dbclient.GetStateCurrent(name)
	if err != nil {
		log.Err(err).Msg("error calling api")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(state)
	if err != nil {
		log.Err(err).Msg("error calling api")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
func StateHistoricalHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	name := chi.URLParam(r, "state")
	dbclient := db.New()
	items, err := dbclient.GetStateHistorical(name)
	if err != nil {
		log.Err(err).Msg("error calling api")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(items)
	if err != nil {
		log.Err(err).Msg("error calling api")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)

}

func USHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("soon"))
}
func USHistoricalHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("soon"))
}
