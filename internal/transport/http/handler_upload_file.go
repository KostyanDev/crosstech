package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/internal/transport/converters"
)

// UploadJsonFile handles the uploading and processing of a JSON file.
//
// @Summary Upload and process a JSON file
// @Description This endpoint receives a JSON file name, processes its content, and inserts relevant data into the database.
// @Tags files
// @Accept  json
// @Produce  json
// @Param file body converters.UploadFileRequest true "File Upload Request"
// @Success 200 {object} map[string]string "status: success, message: File processed successfully"
// @Failure 400 {object} map[string]string "Bad Request: Invalid JSON"
// @Failure 500 {object} map[string]string "Internal Server Error: Failed to process the file"
// @Router /upload_file [post]
func (h *Handler) UploadJsonFile(w http.ResponseWriter, r *http.Request) {
	var request converters.UploadFileRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error("Error decoding request: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.ProcessJSONFile(r.Context(), converters.ToDomainUploadFileName(request))
	if err != nil {
		h.log.Error("Error processing JSON file: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"status": "success", "message": "File processed successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.log.Error(fmt.Sprintf("Write error: %s", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
