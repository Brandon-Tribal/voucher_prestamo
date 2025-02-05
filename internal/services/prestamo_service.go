package services

import (
	"core-migracion/internal/models"
	"fmt"
	"math/rand"
	"time"
)

type prestamoProcessor interface {
	ProcesarSolicitud(solicitud models.SolicitudPrestamo) (models.Voucher, error)
}

type PrestamoService struct {
	Validator Validator
}

func NewPrestamoService(validator Validator) *PrestamoService {
	return &PrestamoService{
		Validator: validator,
	}
}

func (ls *PrestamoService) ProcesarSolicitud(solicitud models.SolicitudPrestamo) (models.Voucher, error) {
	if err := ls.Validator.ValidarSolicitud(solicitud); err != nil {
		return models.Voucher{}, err
	}

	err := ls.ConsumirServicioExterno(solicitud.DatosCliente)
	if err != nil {
		return models.Voucher{}, fmt.Errorf("ha ocurrido un error")
	}

	dpi := ls.ValidacionDPI(solicitud.DatosCliente)
	if dpi != nil {
		return models.Voucher{}, fmt.Errorf("DPI incorrecto")
	}

	selphi := ls.ValidacionSelphi(solicitud.DatosCliente)
	if selphi != nil {
		return models.Voucher{}, fmt.Errorf("validacion con selphi incorrecta")
	}

	totalIngresos := solicitud.DatosLaborales.IngresosMensuales + solicitud.DatosLaborales.OtrosIngresos
	totalEgresos := solicitud.DatosEconomicos.EgresosMensuales

	numeroGestion := generarNumeroGestion()

	voucher := models.Voucher{
		NumeroGestion:   numeroGestion,
		DatosPersonales: solicitud.DatosPersonales,
		TotalIngresos:   int(totalIngresos),
		TotalEgresos:    int(totalEgresos),
	}

	return voucher, nil
}

func generarNumeroGestion() string {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	randomNumber := rng.Intn(1000000)

	return fmt.Sprintf("GEST-%d", randomNumber)
}

func (ps *PrestamoService) ConsumirServicioExterno(datosCliente models.DatosCliente) error {
	//err := errors.New("error inesperado en DPI")
	//log.Println(err)
	//return err
	return nil
}

func (ps *PrestamoService) ValidacionDPI(dpi models.DatosCliente) error {
	//err := errors.New("error en validacion del dpi")
	//return err
	return nil
}

func (ps *PrestamoService) ValidacionSelphi(selphi models.DatosCliente) error {
	//err := errors.New("error en validacion de selphi")
	//return err
	return nil
}
