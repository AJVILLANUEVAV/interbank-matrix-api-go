# Interbank Matrix API

API REST de matrices con Go y Fiber. Valida matrices rectangulares, aplica rotación de 90 grados y calcula factorización QR.

## Estado

Scaffold inicial. La implementación, pruebas y Dockerfile se agregarán en el siguiente paso.

## Variables de entorno

- `PORT`: puerto HTTP, por defecto `8080`.
- `STATISTICS_API_URL`: URL de la API de estadísticas.
- `JWT_SECRET`: secreto JWT, requerido cuando se habilita autenticación.
