package handlers

import (
	"net/http"
)

func StateHandler(w http.ResponseWriter, r *http.Request) {

}
func StateHistoricalHandler(w http.ResponseWriter, r *http.Request) {}

func USHandler(w http.ResponseWriter, r *http.Request) {
	//endpoint /api/v1/us/current.json

}
func USHistoricalHandler(w http.ResponseWriter, r *http.Request) {
	//endpoint /api/v1/us/daily.json
}
