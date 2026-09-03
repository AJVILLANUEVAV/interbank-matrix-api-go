package handler

import (
	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/adapters"
	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/application"
	"github.com/gofiber/fiber/v2"
)

func NewApp(statisticsURL string) *fiber.App {
	statisticsClient := adapters.NewStatisticsHTTPClient(statisticsURL)
	processMatrix := application.NewProcessMatrix(statisticsClient)
	return adapters.NewMatrixHTTPHandler(processMatrix).App()
}
