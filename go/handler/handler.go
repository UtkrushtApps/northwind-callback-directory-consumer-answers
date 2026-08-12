package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"northwind/callbackdirectory/service"
)

type Handler struct {
	svc *service.DirectoryService
}

type phoneNumbersResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func NewHandler(svc *service.DirectoryService) *Handler {
	return &Handler{svc: svc}
}

func writeJSON(w http.ResponseWriter, status int, body phoneNumbersResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func (h *Handler) PhoneNumbers(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	phone := r.URL.Query().Get("phone")

	if strings.TrimSpace(country) == "" || strings.TrimSpace(phone) == "" {
		writeJSON(w, http.StatusBadRequest, phoneNumbersResponse{Error: "country and phone are required"})
		return
	}

	resolved, err := h.svc.GetPhoneNumbers(country, phone)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, phoneNumbersResponse{Error: err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, phoneNumbersResponse{Result: resolved})
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/phone-numbers", h.PhoneNumbers)
	return mux
}
