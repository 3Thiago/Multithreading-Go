package handler

import (
	"encoding/json"
	"net/http"

	"github.com/3Thiago/Multithreading-Go/internal/infra/service"
	"github.com/go-chi/chi"
)

type ApiHandler struct {
	ServiceInfra *service.ServiceInfra
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{
		ServiceInfra: service.NewServiceInfra(),
	}
}

func (h *ApiHandler)GetCepHandler(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if len(cep) < 8 || len(cep) > 9 {
		http.Error(w, "example: /consulta/28922-340", http.StatusBadRequest)
		return
	}
	resp, err := h.ServiceInfra. GetCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
