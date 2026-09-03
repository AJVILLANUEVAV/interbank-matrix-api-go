# Interbank Matrix API

API REST de matrices con Go y Fiber. Valida matrices rectangulares, aplica rotación de 90 grados y calcula factorización QR.

La estructura sigue arquitectura hexagonal: `internal/matrix` contiene el dominio matemático, `internal/application` el caso de uso, `internal/ports` los contratos y `internal/adapters` los adaptadores HTTP.

## Ejecutar localmente

```bash
go run ./cmd/api
```

Pruebas:

```bash
go test ./...
```

Endpoint principal: `POST /v1/matrices/qr` con `{ "matrix": [[1, 2], [3, 4]] }`.

## API-first

El contrato OpenAPI está en `openapi/openapi.yaml`. Puedes importarlo en [Swagger Editor](https://editor.swagger.io/) o visualizarlo con cualquier herramienta compatible con OpenAPI 3.

## Variables de entorno

- `PORT`: puerto HTTP, por defecto `8080`.
- `STATISTICS_API_URL`: URL de la API de estadísticas.
- `JWT_SECRET`: secreto JWT, requerido cuando se habilita autenticación.
- `WEB_ORIGIN`: origen del frontend permitido por CORS, por ejemplo `https://interbank-matrix-qzaru4ygq-ali-vp.vercel.app`.

En Render debes configurar `WEB_ORIGIN=https://interbank-matrix-qzaru4ygq-ali-vp.vercel.app` y hacer redeploy.

Cuando `JWT_SECRET` está configurado, `POST /v1/matrices/qr` requiere `Authorization: Bearer <token>`. Matrix API firma un token de servicio separado para llamar a Statistics API.
