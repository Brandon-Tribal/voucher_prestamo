package main

import (
	"core-migracion/api/controllers"
	"core-migracion/internal/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	validator := &services.ValidationService{
		EnableOtrosIngresos: true,
	}
	prestamoService := &services.PrestamoService{
		Validator: validator,
	}

	prestamoController := controllers.NewPrestamoController(prestamoService)

	r := mux.NewRouter()
	r.HandleFunc("/api/solicitudes", prestamoController.ProcesarSolicitudHandler).Methods("POST")

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
