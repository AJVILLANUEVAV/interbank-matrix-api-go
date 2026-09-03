FROM golang:1.27-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /matrix-api ./cmd/api

FROM gcr.io/distroless/static-debian12
COPY --from=build /matrix-api /matrix-api
EXPOSE 8080
ENTRYPOINT ["/matrix-api"]
