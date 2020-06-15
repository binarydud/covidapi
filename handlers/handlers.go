package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/binarydud/covidapi/db"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/hlog"
)

type dBKeyID string

const requestDBKey dBKeyID = "client"

// DBMiddleware ...
func DBMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		dbclient := db.New()

		ctx = context.WithValue(ctx, requestDBKey, dbclient)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
func dbFromRequest(r *http.Request) *db.DB {
	ctx := r.Context()
	if client, ok := ctx.Value(requestDBKey).(*db.DB); ok {
		return client
	}
	return nil
}

// HealthHandler ...
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

// StateHandler ...
func StateHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	dbclient := dbFromRequest(r)
	name := chi.URLParam(r, "state")
	name = strings.ToUpper(name)
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

// StateHistoricalHandler ...
func StateHistoricalHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	dbclient := dbFromRequest(r)
	name := chi.URLParam(r, "state")
	name = strings.ToUpper(name)
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

// USHandler ...
func USHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	dbclient := dbFromRequest(r)
	log.Info().Msg("Getting current US data")
	current, err := dbclient.GetUSCurrent()
	if err != nil {
		log.Err(err).Msg("error getting current record for US")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(current)
	if err != nil {
		log.Err(err).Msg("error encoding records")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

// USHistoricalHandler ...
func USHistoricalHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	dbclient := dbFromRequest(r)
	log.Info().Msg("Getting historical US data")
	values, err := dbclient.GetUSHistorical()
	if err != nil {
		log.Err(err).Msg("error getting db records for US")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(values)
	if err != nil {
		log.Err(err).Msg("error encoding records")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}

// StatesDailyHandler ...
func StatesDailyHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	dbclient := dbFromRequest(r)
	states, err := dbclient.GetStatesDaily()
	if err != nil {
		log.Err(err).Msg("error getting states records")
	}
	body, err := json.Marshal(states)
	if err != nil {
		log.Err(err).Msg("error encoding records")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
