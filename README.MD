## 1️⃣ Descripción General
Este documento describe el flujo de la solicitud de préstamo personal en el sistema. Se detallan los pasos desde la recepción de la solicitud hasta la generación del voucher, incluyendo validaciones y reglas de negocio.

## 2️⃣ Flujo del Sistema

| Paso                 | Acción                                   | Posible Resultado |
|----------------------|------------------------------------------|-------------------|
| 📌 Cliente           | Envío de JSON con datos                  | Solicitud recibida |
| ✅ Validaciones      | Validar estructura del JSON              | OK / Error 400 |
| 🛠️ Reglas de Negocio | Evaluar ingresos, egresos, plazo         | Aprobado / Rechazado |
| 🎟️ Generar Voucher   | Si se aprueba, asignar número de gestión | Voucher generado |
| ❌ Error             | Si falla, enviar respuesta de rechazo    | `"Ha ocurrido un error"` |


## 3️⃣ Pasos del Flujo

1. **El cliente envía una solicitud** (`POST /api/solicitudes`) con los datos personales, laborales y del préstamo en formato JSON.
2. **El sistema valida la estructura del JSON** para verificar que los datos sean correctos en formato, longitudes y tipos de datos.
3. **Se aplican las reglas de negocio**, que incluyen:
   - Ingresos deben ser mayores o iguales a 4500 GTQ.
   - Egresos no pueden ser mayores o iguales al 40% de los ingresos.
   - El monto solicitado debe estar entre 10,000 y 160,000 GTQ.
   - El plazo debe ser de 6, 9, 12, 18, 24, 36 o 48 meses.
   - La antigüedad laboral debe ser de al menos 6 meses y 7 días.
   - Validaciones externas: DPI y Selphi.
4. **Si la solicitud es aprobada**, se genera un número de gestión aleatorio y un voucher con la información consolidada.
5. **Si la solicitud no es aprobada**, el sistema devuelve un mensaje indicando el motivo del rechazo.
6. **Si un servicio externo falla**, el sistema devuelve un mensaje de error genérico: `"Ha ocurrido un error"`.

## 4️⃣ Respuestas del Sistema

### ✅ Solicitud Aprobada
```json
{
  "numero_gestion": "GEST-123456",
  "datos_personales": { "nombres": "Juan", "apellidos": "Pérez", "correo": "juan@example.com" },
  "total_ingresos": 5000,
  "total_egresos": 2000
}
```

### ❌ Solicitud Rechazada por Reglas de Negocio
```json
{
  "error": "No aplica para el préstamo: ingresos menores a 4500."
}
```

### ⚠️ Error en un Servicio Externo
```json
{
  "error": "Ha ocurrido un error"
}
```

## 5️⃣ Consideraciones Especiales

- **El campo "Otros Ingresos" depende de una Feature Flag**.
  - Si `EnableOtrosIngresos = false`, cualquier valor distinto de cero en `otros_ingresos` causa un error.
  - Si `EnableOtrosIngresos = true`, se permite su uso y validaciones.

## 6️⃣ Cómo Ejecutar la API

### 🔹 Prerrequisitos
- Tener Go instalado.
- Clonar el repositorio.

### 🔹 Pasos para Ejecutar
```bash
git clone https://github.com/Brandon-Tribal/voucher_prestamo.git
cd core-migracion
go mod tidy
go run main.go
```

- La API estará disponible en: `http://localhost:8080/api/solicitudes`

---

### 🚀 Conclusión
Este documento describe el flujo del sistema de solicitud de préstamos. Se detallan los pasos, reglas de negocio y respuestas esperadas, permitiendo a los desarrolladores y analistas entender el funcionamiento general del servicio.

