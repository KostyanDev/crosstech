package http

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "app/docs"
)

func RegisterRoutes(router *mux.Router, handler *Handler) {
	router.HandleFunc("/upload_file", handler.UploadJsonFile).Methods("POST")

	router.HandleFunc("/signals", handler.GetSignals).Methods("GET")
	router.HandleFunc("/signal/create", handler.CreateSignal).Methods("POST")
	router.HandleFunc("/signal/update", handler.UpdateSignal).Methods("POST")

	router.HandleFunc("/tracks", handler.GetTracks).Methods("GET")
	router.HandleFunc("/track/create", handler.CreateTrack).Methods("POST")
	router.HandleFunc("/track/update", handler.UpdateTrack).Methods("POST")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
