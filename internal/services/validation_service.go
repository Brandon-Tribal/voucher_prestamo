package services

import (
	"core-migracion/internal/models"
	"core-migracion/internal/utils"
	"errors"
	"fmt"
	"time"
)

type Validator interface {
	ValidarSolicitud(solicitud models.SolicitudPrestamo) error
}

type ValidationService struct {
	EnableOtrosIngresos bool
}

func (v *ValidationService) ValidarSolicitud(solicitud models.SolicitudPrestamo) error {
	//valida la estructura del modelo
	if err := utils.ValidarEstructura(solicitud); err != nil {
		return err
	}

	return validarReglasNegocio(solicitud, v.EnableOtrosIngresos)
}

func validarReglasNegocio(solicitud models.SolicitudPrestamo, EnableOtrosIngresos bool) error {
	if solicitud.DatosLaborales.IngresosMensuales < 4500 {
		return errors.New("ingresos menores a GTQ 4,500 no aplican")
	}

	if solicitud.DatosEconomicos.EgresosMensuales >= solicitud.DatosLaborales.IngresosMensuales {
		return errors.New("egresos no pueden superar o igual los ingresos")
	}
	if solicitud.DatosEconomicos.EgresosMensuales >= 0.4*solicitud.DatosLaborales.IngresosMensuales {
		return errors.New("egresos no pueden superar el 40% de los ingresos")
	}

	fechaLimite := time.Now().AddDate(0, -6, -7)
	if solicitud.DatosLaborales.FechaInicio.After(fechaLimite) {
		return errors.New("antigüedad laboral minima es 6 meses y 7 dias")
	}

	if solicitud.DatosPrestamo.Monto < 10000 || solicitud.DatosPrestamo.Monto > 160000 {
		return errors.New("el monto debe estar entre GTQ 10,000 y 160000")
	}
	if !esPlazoValido(solicitud.DatosPrestamo.Plazo) {
		return errors.New("el plazo no es valido")
	}

	if !solicitud.DatosCliente.ValidacionDPI {
		return errors.New("validacion de dpi fallida")
	}
	if !solicitud.DatosCliente.ValidacionSelphi {
		return errors.New("validacion de Selphi fallida")
	}
	if len(solicitud.DatosPersonales.Telefono) == 8 {
		primerDigito := solicitud.DatosPersonales.Telefono[0]
		if primerDigito != '5' && primerDigito != '6' && primerDigito != '7' {
			return fmt.Errorf("el teléfono debe comenzar con 5, 6 o 7")
		}
	}

	if !EnableOtrosIngresos && solicitud.DatosLaborales.OtrosIngresos > 0 {
		return errors.New("el campo 'Otros Ingresos' no esta habilitado actualmente")
	}

	return nil
}

func esPlazoValido(plazo int) bool {
	plazosPermitidos := []int{6, 9, 12, 18, 24, 36, 48}
	for _, p := range plazosPermitidos {
		if p == plazo {
			return true
		}
	}
	return false
}
