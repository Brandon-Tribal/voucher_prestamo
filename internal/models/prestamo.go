package models

import "time"

type DatosPersonales struct {
	Nombres      string `json:"nombres" validate:"required,max=24"`
	Apellidos    string `json:"apellidos" validate:"required,max=24"`
	Correo       string `json:"correo" validate:"required,email,max=64"`
	Telefono     string `json:"telefono" validate:"required,len=8,numeric"`
	Departamento string `json:"departamento" validate:"required"`
	Municipio    string `json:"municipio" validate:"required"`
}

type DatosLaborales struct {
	NombreEmpresa      string    `json:"nombre_empresa" validate:"required,max=32"`
	ActividadEconomica string    `json:"actividad_economica" validate:"required"`
	FechaInicio        time.Time `json:"fecha_inicio" validate:"required"`
	IngresosMensuales  float64   `json:"ingresos_mensuales" validate:"required,min=0"`
	OtrosIngresos      float64   `json:"otros_ingresos" validate:"min=0"`
}

type DatosEconomicos struct {
	NumeroDependientes int     `json:"numero_dependientes" validate:"required,min=0"`
	EgresosMensuales   float64 `json:"egresos_mensuales" validate:"required,min=0"`
}

type DatosPrestamo struct {
	Monto float64 `json:"monto" validate:"required"`
	Plazo int     `json:"plazo" validate:"required"`
}

type DatosCliente struct {
	ValidacionDPI    bool `json:"validacion_dpi"`
	ValidacionSelphi bool `json:"validacion_selphi"`
}

type SolicitudPrestamo struct {
	DatosPersonales DatosPersonales `json:"datos_personales"`
	DatosLaborales  DatosLaborales  `json:"datos_laborales"`
	DatosEconomicos DatosEconomicos `json:"datos_economicos"`
	DatosPrestamo   DatosPrestamo   `json:"datos_prestamo"`
	DatosCliente    DatosCliente    `json:"datos_cliente"`
}

type Voucher struct {
	NumeroGestion   string          `json:"numero_gestion"`
	DatosPersonales DatosPersonales `json:"datos_personales"`
	TotalIngresos   int             `json:"total_ingresos"`
	TotalEgresos    int             `json:"total_egresos"`
}
