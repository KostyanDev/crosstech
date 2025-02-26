package http

import (
	_ "app/internal/domain"
	"app/internal/transport/converters"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateSignal
//
// @Summary Create a signal
// @Description Creates a new signal in the system
// @Tags signals
// @Accept  json
// @Produce  json
// @Param request body converters.CreatSignalRequest true "Signal creation data"
// @Success 200 {object} map[string]string "Successful signal creation"
// @Failure 400 {object} map[string]string "Invalid request data"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /signal/create [post]
func (h *Handler) CreateSignal(w http.ResponseWriter, r *http.Request) {
	var request converters.CreatSignalRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error("Error decoding request: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.CreateSignal(r.Context(), converters.ToDomainCreatSignal(request))
	if err != nil {
		h.log.Error("Error creating signal: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "success", "message": "Signal created successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error(fmt.Sprintf("Write error: %s", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateSignal
//
// @Summary Update a signal
// @Description Updates an existing signal in the system
// @Tags signals
// @Accept  json
// @Produce  json
// @Param request body converters.UpdateSignalRequest true "Signal update data"
// @Success 200 {object} map[string]string "Successful signal update"
// @Failure 400 {object} map[string]string "Invalid request data"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /signal/update [post]
func (h *Handler) UpdateSignal(w http.ResponseWriter, r *http.Request) {
	var request converters.UpdateSignalRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error("Error decoding request: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.UpdateSignalByParam(r.Context(), converters.ToDomainUpdateSignal(request))
	if err != nil {
		h.log.Error("Error updating signal: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "success", "message": "Signal updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error(fmt.Sprintf("Write error: %s", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetSignals
//
// @Summary Get signals
// @Description Retrieves signals based on query parameters
// @Tags signals
// @Accept  json
// @Produce  json
// @Param signal_id query int false "Filter by signal ID"
// @Param track_id query int false "Filter by track ID"
// @Success 200 {array} domain.Signal "List of signals"
// @Failure 400 {object} map[string]string "Invalid query parameters"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /signal [get]
func (h *Handler) GetSignals(w http.ResponseWriter, r *http.Request) {
	request, err := converters.ParseGetSignalRequest(r.URL.Query())
	if err != nil {
		http.Error(w, "invalid query parameters", http.StatusBadRequest)
		return
	}

	signals, err := h.service.GetSignalByParam(r.Context(), request)
	if err != nil {
		h.log.Error("Error fetching signals: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(converters.ToRespSignals(signals)); err != nil {
		h.log.Error("Write error: ", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
