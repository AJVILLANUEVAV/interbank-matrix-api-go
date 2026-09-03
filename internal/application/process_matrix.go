package application

import (
	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/matrix"
	"github.com/AJVILLANUEVAV/interbank-matrix-api-go/internal/ports"
)

type ProcessMatrix struct {
	statistics ports.StatisticsGateway
}

func NewProcessMatrix(statistics ports.StatisticsGateway) *ProcessMatrix {
	return &ProcessMatrix{statistics: statistics}
}

func (useCase *ProcessMatrix) Execute(input [][]float64) (ports.MatrixResult, error) {
	if err := matrix.Validate(input); err != nil {
		return ports.MatrixResult{}, err
	}
	rotated := matrix.RotateClockwise(input)
	q, r, err := matrix.QR(rotated)
	if err != nil {
		return ports.MatrixResult{}, err
	}
	statistics, err := useCase.statistics.Calculate(rotated, q, r)
	if err != nil {
		return ports.MatrixResult{}, err
	}
	return ports.MatrixResult{RotatedMatrix: rotated, Q: q, R: r, Statistics: statistics}, nil
}
