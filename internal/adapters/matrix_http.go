package adapters

import (
	"os"

	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/application"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type matrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type MatrixHTTPHandler struct {
	process *application.ProcessMatrix
}

func NewMatrixHTTPHandler(process *application.ProcessMatrix) *MatrixHTTPHandler {
	return &MatrixHTTPHandler{process: process}
}

func (handler *MatrixHTTPHandler) App() *fiber.App {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(cors.New(cors.Config{AllowOrigins: "http://localhost:5173,http://localhost:3000"}))
	app.Get("/health", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status": "ok"}) })
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
