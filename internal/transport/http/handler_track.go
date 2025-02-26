package http

import (
	_ "app/internal/domain"
	"app/internal/transport/converters"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateTrack creates a new track.
//
// @Summary Create a track
// @Description This endpoint allows the creation of a new track in the system.
// @Tags tracks
// @Accept  json
// @Produce  json
// @Param track body converters.CreatTrackRequest true "Track Creation Request"
// @Success 201 {object} map[string]string "status: success, message: Track created successfully"
// @Failure 400 {object} map[string]string "Bad Request: Invalid JSON"
// @Failure 500 {object} map[string]string "Internal Server Error: Failed to create track"
// @Router /track/create [post]
func (h *Handler) CreateTrack(w http.ResponseWriter, r *http.Request) {
	var request converters.CreatTrackRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error("Error decoding request: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.CreateTrack(r.Context(), converters.ToDomainCreatTrack(request))
	if err != nil {
		h.log.Error("Error creating track: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "success", "message": "Track created successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 вместо 200 для создания
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error(fmt.Sprintf("Write error: %s", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// UpdateTrack updates an existing track.
//
// @Summary Update a track
// @Description This endpoint updates an existing track with new values.
// @Tags tracks
// @Accept  json
// @Produce  json
// @Param track body converters.UpdateTrackRequest true "Track Update Request"
// @Success 200 {object} map[string]string "status: success, message: Track updated successfully"
// @Failure 400 {object} map[string]string "Bad Request: Invalid JSON"
// @Failure 404 {object} map[string]string "Not Found: Track not found"
// @Failure 500 {object} map[string]string "Internal Server Error: Failed to update track"
// @Router /track/update [put]
func (h *Handler) UpdateTrack(w http.ResponseWriter, r *http.Request) {
	var request converters.UpdateTrackRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error("Error decoding request: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.UpdateTrackByParam(r.Context(), converters.ToDomainUpdateTrack(request))
	if err != nil {
		h.log.Error("Error updating track: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "success", "message": "Track updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error(fmt.Sprintf("Write error: %s", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetTracks retrieves tracks based on query parameters.
//
// @Summary Retrieve tracks
// @Description This endpoint fetches tracks based on optional filters such as ID, source, or target.
// @Tags tracks
// @Accept  json
// @Produce  json
// @Param track_id query int false "Track ID"
// @Param source query string false "Source station of the track"
// @Param target query string false "Target station of the track"
// @Success 200 {array} domain.Track "List of tracks"
// @Failure 400 {object} map[string]string "Bad Request: Invalid query parameters"
// @Failure 404 {object} map[string]string "Not Found: No tracks found"
// @Failure 500 {object} map[string]string "Internal Server Error: Failed to retrieve tracks"
// @Router /track [get]
func (h *Handler) GetTracks(w http.ResponseWriter, r *http.Request) {
	request, err := converters.ParseGetTrackRequest(r.URL.Query())
	if err != nil {
		http.Error(w, "invalid query parameters", http.StatusBadRequest)
		return
	}

	tracks, err := h.service.GetTrackByParam(r.Context(), request)
	if err != nil {
		h.log.Error("Error fetching tracks: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if len(tracks) == 0 {
		http.Error(w, "no tracks found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(converters.ToRespTracks(tracks)); err != nil {
		h.log.Error("Write error: ", err)
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
