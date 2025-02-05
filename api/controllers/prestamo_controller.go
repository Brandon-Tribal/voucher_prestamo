package controllers

import (
	"core-migracion/internal/models"
	"core-migracion/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type PrestamoController struct {
	PrestamoService *services.PrestamoService
}

func NewPrestamoController(ps *services.PrestamoService) *PrestamoController {
	return &PrestamoController{PrestamoService: ps}
}

func (lc *PrestamoController) ProcesarSolicitudHandler(w http.ResponseWriter, r *http.Request) {
	var solicitud models.SolicitudPrestamo
	err := json.NewDecoder(r.Body).Decode(&solicitud)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer la solicitud: %v", err), http.StatusBadRequest)
		return
	}

	voucher, err := lc.PrestamoService.ProcesarSolicitud(solicitud)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error en la solicitud de prestamo: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(voucher)
}
