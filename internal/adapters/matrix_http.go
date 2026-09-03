package adapters

import (
	"errors"
	"os"

	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/application"
	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type matrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type MatrixHTTPHandler struct {
	process   *application.ProcessMatrix
	jwtSecret string
}

func NewMatrixHTTPHandler(process *application.ProcessMatrix, jwtSecret string) *MatrixHTTPHandler {
	return &MatrixHTTPHandler{process: process, jwtSecret: jwtSecret}
}

func (handler *MatrixHTTPHandler) App() *fiber.App {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(cors.New(cors.Config{AllowOrigins: WebOrigin()}))
	app.Get("/health", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) })
	if handler.jwtSecret != "" {
		app.Use("/v1", JWTMiddleware(handler.jwtSecret))
	}
	app.Post("/v1/matrices/qr", handler.processMatrix)
	return app
}

func (handler *MatrixHTTPHandler) processMatrix(c *fiber.Ctx) error {
	var request matrixRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON body"})
	}
	result, err := handler.process.Execute(request.Matrix)
	if err != nil {
		if errors.Is(err, ports.ErrStatisticsUnavailable) {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func StatisticsURL() string {
	if url := os.Getenv("STATISTICS_API_URL"); url != "" {
		return url
	}
	return "http://localhost:8081"
}

func JWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func WebOrigin() string {
	if origin := os.Getenv("WEB_ORIGIN"); origin != "" {
		return "http://localhost:5173,http://localhost:3000," + origin
	}
	return "http://localhost:5173,http://localhost:3000"
}
